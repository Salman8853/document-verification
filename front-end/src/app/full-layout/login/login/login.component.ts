import { Component, OnInit, ViewChild } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { TextBoxComponent } from '@progress/kendo-angular-inputs';
import { NotificationService } from '@progress/kendo-angular-notification';
import { tap } from 'rxjs/operators';
import { LoginService } from 'src/app/services/auth/login.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
})
export class LoginComponent implements OnInit {
  loginForm = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.email]),
    password: new FormControl(''),
  });

  public userdata = { email: '', password: '' };
  public userobj: any;

  public showPassword: boolean = true;
  @ViewChild('password') textbox!: TextBoxComponent;
  public ngAfterViewInit(): void {
    this.textbox.input.nativeElement.type = 'password';
  }

  constructor(
    private _loginServ: LoginService,
    private router: Router,
    private notificationService: NotificationService
  ) {}

  ngOnInit(): void {
    const userCred = this.getSavedUserName();
    if (userCred) {
      const email = JSON.parse(userCred).email;
      const password = JSON.parse(userCred).password;
      if (email) {
        this.loginForm.patchValue({
          email,
          password,
        });
      }
    }
  }

  public toggleVisibility(): void {
    const inputEl = this.textbox.input.nativeElement;
    inputEl.type = inputEl.type === 'password' ? 'text' : 'password';
    this.showPassword = !this.showPassword;
  }

  onSubmit() {
    console.log(this.loginForm.value);
    this.login();
  }

  getSavedUserName() {
    return localStorage.getItem('savedUserCredentials');
  }

  public login(): void {
    this.loginForm.markAllAsTouched();

    this.userdata = {
      email: this.loginForm.controls['email'].value? this.loginForm.controls['email'].value : '',
      password: this.loginForm.controls['password'].value? this.loginForm.controls['password'].value: '',
    };

    this._loginServ.login(this.userdata).subscribe(user=>{
        this.userobj = user
        console.log(this.userobj.roles);
        if(this.userobj.roles[0] === "ROLE_ADMIN" || this.userobj.roles[0] === "ROLE_SUBADMIN"){
          this._loginServ.loggedInUserDetails.next(this.userobj);
          this.router.navigate(['/home']);
        }else{
          this.router.navigate(['/login'])
        }
      },err=>{
        this.showError(err.error.message)
      });
  }

  public showError(errtext: string): void {
    this.notificationService.show({
      content: errtext,
      hideAfter: 3000,
      position: { horizontal: 'center', vertical: 'bottom' },
      animation: { type: 'fade', duration: 300 },
      type: { style: 'error', icon: true },
    });
  }
  public showSuccess(successtext: string): void {
    this.notificationService.show({
      content: successtext,
      hideAfter: 3000,
      position: { horizontal: 'center', vertical: 'bottom' },
      animation: { type: 'fade', duration: 300 },
      type: { style: 'success', icon: true },
    });
  }

  saveUser(value: any) {
    const remembered = value.target.checked;
    if (remembered) {
      localStorage.setItem(
        'savedUserCredentials',
        JSON.stringify(this.loginForm.value)
      );
    }
  }
}
