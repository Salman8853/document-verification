import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'fadv-DQF-NG';
  islogoutAlert = false;
  countdown:any;
  interval:any;
  constructor(private router: Router) { }
  ngOnInit(){
    console.clear();
    //setLocalStorage();
    this.inactivityTime();
  }
  inactivityTime() {
    let time : any;
    //this.countdown = 30;
    const setTimer = () => {
      if(this.islogoutAlert=== false){
        this.countdown = 30;
      }
      if (sessionStorage.getItem('UserInfo')) {
        
        clearInterval(this.interval);
        this.islogoutAlert = true;
        this.interval = setInterval(() => {
          if (this.countdown > 0) {
            this.countdown--;
            //console.log(this.countdown)
          }
          else {
            this.islogoutAlert = false;
            clearInterval(time);
            this.logout();
          }
        }, 1000)
      }
    }
    function resetTimer() {
      clearInterval(time);
      //setTimeout(logout, 1000 * 60 * 15);
      time = setInterval(setTimer, 1000 * 60 * 29.5); 
    }
    window.onload = resetTimer;
    //window.addEventListener('load', resetTimer, true);
    let events = ['mousedown', 'mousemove', 'keypress', 'scroll', 'touchstart'];
    events.forEach(function (name) {
      document.addEventListener(name, resetTimer, true);
    });
  }

  logout(){
    if (sessionStorage.getItem('UserInfo')) {
      sessionStorage.removeItem('UserInfo');
      sessionStorage.removeItem('userToken');
      sessionStorage.removeItem('userDetails');
      this.router.navigate(['/login']);
      window.location.reload();
    }
    clearInterval(this.interval);
  }

  logoutConfirmationAlert() {
    clearInterval(this.interval);
    this.islogoutAlert = false;
    this.logout();
  }

  stayConfirmationAlert() {
    clearInterval(this.interval);
    this.islogoutAlert = false;
  }
}
