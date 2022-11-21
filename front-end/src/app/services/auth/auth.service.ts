import { Injectable } from '@angular/core';
import {
  ActivatedRouteSnapshot,
  CanActivate,
  Router,
  RouterStateSnapshot,
} from '@angular/router';
import { LoginService } from './login.service';

@Injectable({
  providedIn: 'root',
})
export class AuthGuard implements CanActivate {
  user: any;

  constructor(
    private router: Router,
    public _loginService: LoginService
  ) {}

  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot) {
    //return true
    if (this._loginService.isLoggedIn()) {
      const user = this._loginService.getUserDetails();
      this._loginService.userRole = user.roles[0];

      if (route.data['role'].indexOf(user.roles[0]) === -1) {
        this.router.navigate(['/home']);
        return false;
      }
      return true;
    }

    // not logged in so redirect to login page with the return url
    this.router.navigate(['/login'], { queryParams: { returnUrl: state.url } });
    return false;
  }

  canLoad(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): boolean {
    return true
    if (this._loginService.isLoggedIn()) {
      const user = this._loginService.getUserDetails();

      if (!user) {
        this.router.navigate(['/login']);
      }

      return true;
    }

    // not logged in so redirect to login page with the return url
    // this.router.navigate(['/login'], { queryParams: { returnUrl: state.url } });
    return true;
  }
}
