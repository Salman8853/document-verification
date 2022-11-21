import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable()
export class AppConfig {
  constructor(private httpClient: HttpClient) {}

  authenticationUrl = '';
  env = '';

  ensureInit(): Promise<any> {
    return new Promise((r, e) => {
      this.httpClient.get('./assets/jsons/api-config.json').subscribe(
        (content) => {
          console.log('123--->', content);
          Object.assign(this, content);

          r(this);
        },
        (reason) => e(reason)
      );
    });
  }
}
