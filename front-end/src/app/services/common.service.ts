import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class CommonService {
  sidePanel = new BehaviorSubject(true);
  sideMenuWidth = new BehaviorSubject(0);
  constructor(private _http: HttpClient) { }
  fetchMenuStructre(): Observable<any>{
    return this._http.get(`./assets/jsons/menu.json`);
  }
}
