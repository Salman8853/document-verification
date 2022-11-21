import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { BehaviorSubject, Subscription } from 'rxjs';
import { LoginService } from 'src/app/services/auth/login.service';
import { CommonService } from 'src/app/services/common.service';
import { UiFlowService } from 'src/app/services/ui-flow.service';

@Component({
  selector: 'app-full-layout',
  templateUrl: './full-layout.component.html',
  styleUrls: ['./full-layout.component.css'],
})
export class FullLayoutComponent implements OnInit {
  public size = 'medium';
  public shape = 'rectangle';
  public rounded = 'medium';
  public fillMode = 'solid';
  public themeColor = 'base';
  showMenu = true;

  $loginDetailsSub: Subscription | undefined;
  $sidePanelSub!: Subscription;
  currentUserName = '';
  loggedInUserDetails: any;

  menuData : any = [];

  constructor(private loginobj: LoginService, private router : Router, private uiFlow: UiFlowService, private commonSrv : CommonService) {}

  ngOnInit(): void {
    this.getMenuItems();
    this.$loginDetailsSub = this.loginobj.loggedInUserDetails.subscribe(data=>{
      this.loggedInUserDetails = data;
      console.log(this.loggedInUserDetails);
      if(this.loggedInUserDetails){
        this.currentUserName = `${this.loggedInUserDetails.firstName} ${this.loggedInUserDetails.lastName}`
      }
    })
  }

  showSidebar() {
    this.showMenu = !this.showMenu;
    this.commonSrv.sidePanel.next(this.showMenu)
  }

  navigate(route:any){
    console.log("Router", route)
    this.router.navigate([route])
  }

  onlogout(){
    sessionStorage.removeItem('UserInfo');
    sessionStorage.removeItem('userToken');
    sessionStorage.removeItem('userDetails');
    this.router.navigate(['/login']);
  }

  getMenuItems() {
    this.uiFlow.fetchMenuStructre().subscribe((data) => {
      this.menuData = data.allMenuItems;
    });
  }

  openDialog(menu: { hasOwnProperty: (arg0: string) => any; menuItems: string | any[]; path: any; }) {
    if (!menu.hasOwnProperty('menuItems') || !menu.menuItems.length) {
      console.log(menu.path);
      this.router.navigate([menu.path]);
    }
  }

}
