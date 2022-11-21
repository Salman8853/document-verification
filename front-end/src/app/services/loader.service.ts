import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
@Injectable({
  providedIn: 'root'
})
export class LoaderService {
  public isLoading = new BehaviorSubject(true);
  public isError = new BehaviorSubject({message: "", status: ""});
  public apiLoader = new BehaviorSubject(false);
  constructor() { }

  setLoading(loading: boolean): void {
    if (loading) {
      this.apiLoader.next(true);
    } else {
      this.apiLoader.next(false);
    }
  }
}
