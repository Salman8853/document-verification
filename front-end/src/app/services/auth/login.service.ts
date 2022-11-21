import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable, of, throwError } from 'rxjs';
import { AppConfig } from '../app-config.service';
import { UiFlowService } from '../ui-flow.service';
import jwt_decode from 'jwt-decode';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class LoginService {
  loggedInUserDetails = new BehaviorSubject<any>(null);
  userRole: string = '';
  authURL = '';
  deployedEnv = '';
  loggedinUser: any = null;
  loggedinId = '';

  constructor(
    private http: HttpClient,
    public config: AppConfig,
    private _uiFlow: UiFlowService
  ) {
    this.authURL = config.authenticationUrl;
    this.deployedEnv = config.env;
  }
  isLoggedIn() {
    const loggedinUser = sessionStorage.getItem('user');
    if (loggedinUser) {
      const userToken = JSON.parse(loggedinUser).Value;
      const decodedToken: any = jwt_decode(userToken);
      const expiry = decodedToken.exp;

      if (expiry * 1000 >= Date.now()) {
        const userEmail = decodedToken.userid.toLowerCase();
        return this.fetchUserDetails(userToken).subscribe((data) => {
          if (data.length) {
            sessionStorage.setItem('userDetails', JSON.stringify(data[0]));
            this.loggedInUserDetails.next(Object(data[0]));

            return true;
          }
          return false;
        });
      }
      return false;
    }

    return false;
  }
  getUserObj() {
    return this.loggedInUserDetails;
  }
  getUserDetails() {
    const user = sessionStorage.getItem('user');
    return user ? JSON.parse(user) : '';
  }

  // login(obj: any, url: string): Observable<any> {
  //   const formData = new FormData();
  //   formData.append('userid', obj.email);
  //   formData.append('password', obj.password);
  //   const loginObj = {
  //     userName: obj.username,
  //     password: obj.password,
  //   };
  //   console.clear();
  //   return this.http.post<any>(`${this.authURL}${url}`, formData).pipe(
  //     map((el) => {
  //       console.log(el);
  //       let userToken = '';

  //       if (el.Value) {
  //         userToken = el.Value;
  //         sessionStorage.setItem('user', JSON.stringify(el));
  //         sessionStorage.setItem('userToken', userToken);
  //         return el;
  //       }
  //       return throwError('Invalid Value');
  //     })
  //   );
  // }

  login(obj: any): Observable<any> {
    // const loginObj = {
    //   email: obj.username,
    //   password: obj.password,
    // }
    sessionStorage.removeItem('user');
    const userTemp = {
        "success": true,
        "token": "eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJTYWNoaW4uUHJhYmhha2FyYUBmYWR2LmNvbSIsImlhdCI6MTYyNTEyMTQ4NywiZXhwIjoxNjI1MjA3ODg3fQ.WsOD35ELEBa1_uFdF-dvTHwUwymRSyFle7dAAP1ky9_3FZerV339kF2U_C9H8CRzL6vAiTrkM6TyELsIppzOmA",
        "type": "Bearer",
        "id": 1,
        "username": "Sachin.Prabhakara@fadv.com",
        "email": "Sachin.Prabhakara@fadv.com",
        "firstName": "Sachin",
        "lastName": "Prabhakara",
        "roles": [
            "ROLE_SUBADMIN"
        ]
    }
    sessionStorage.setItem('user', JSON.stringify(userTemp));
    return of(userTemp);

    return this.http.post<any>(`${this.authURL}auth/signin`,obj).pipe(map(user=>{
    sessionStorage.setItem('user', JSON.stringify(user))
    return user;

   }));

  }

  fetchUserDetails(value: string): Observable<any> {
    const decodedToken: any = jwt_decode(value);

    const obj = {
      View: 'userDetails',
      Filter: {
        emailId: decodedToken.userid.toLowerCase(),
      },
    };

    // return this.http.post<any>(`${this.apiUrl}workflow/query`, obj).pipe(
    return this.http.get<any>(`./assets/jsons/user.json`).pipe(
      map((ev) => {
        console.log('user-------->', ev);
        if (ev.length) {
          this._uiFlow.loggedInUserDetails.next(Object(ev[0]));
          this.loggedinUser = ev[0];
          this._uiFlow.loggedinUser.next(ev[0].fullName);
          sessionStorage.setItem('ownerID', ev[0].adminInfo.ownerId);
          sessionStorage.setItem('domainName', ev[0].adminInfo.domain);
          sessionStorage.setItem('winESpUserId', ev[0].adminInfo.winESpUserId);
          sessionStorage.setItem('agentFullName', ev[0].fullName);
          sessionStorage.setItem('userDetails', JSON.stringify(ev[0]));
          return ev;
        } else {
          return { message: 'NOUSERFOUND' };
        }
      })
    );
  }



  logout() {
    if (this.loggedinId) {
      this.resetUser();
    } else {
      this.resetUser();
    }
  }
  resetUser() {
    sessionStorage.removeItem('user');
    sessionStorage.removeItem('userToken');
    sessionStorage.removeItem('ownerID');
    sessionStorage.removeItem('ContactSummaryCount');
    sessionStorage.removeItem('domainName');
    sessionStorage.removeItem('userDetails');
    window.location.reload();
  }


}
