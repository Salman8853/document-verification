import { ChangeDetectorRef, Component, OnInit } from '@angular/core';
import { LoaderService } from 'src/app/services/loader.service';

@Component({
  selector: 'app-loader',
  templateUrl: './loader.component.html',
  styleUrls: ['./loader.component.css']
})
export class LoaderComponent implements OnInit {
  loading = false;
  showError = false;
  errorStatus ={message: "", status: ""}

  constructor(private loaderService: LoaderService,private cd: ChangeDetectorRef) {}
 
  ngOnInit() {
    // this.loaderService.isLoading.subscribe((v) => {
    //   this.loading = v;
    // });
    this.loaderService.apiLoader.subscribe((v:any) => {
      // console.log("loader value ",v);
      this.loading = v;
      this.cd.detectChanges();
    });
    this.loaderService.isError.subscribe((e:any)=>{
      this.errorStatus = {...this.errorStatus, message: e.message, status: e.status};
      this.onError();
      console.log(e);
    })
    
  }

  onError(){
    this.showError= this.errorStatus.message !== "";
  }
  
  closeErrorDialog(status:any) {
    if (status === 'yes') {
      this.showError = false;
    } else {
      this.showError = false;
    }
  }

}
