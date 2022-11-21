import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class UiFlowService {
  loggedInUserDetails = new BehaviorSubject<any>(null);
  loggedinUser = new BehaviorSubject('');
  constructor(private _http: HttpClient) {}
  fetchMenuStructre(): Observable<any>{
    return this._http.get(`./assets/jsons/menu.json`);
  }
}
