import { Injectable } from '@angular/core';
import { Wallet } from '../app.datatypes';
import { Observable } from 'rxjs/Observable';
import { HttpClient } from '@angular/common/http';

@Injectable()
export class ApiService {

  private url = 'http://178.62.235.177/api/';

  constructor(
    private httpClient: HttpClient,
  ) { }

  getWalletsList(): Observable<Wallet[]> {
    return this.get('wallets/list').map(response => response.wallets);
  }

  private get(url, params = null, options = {}) {
    return this.httpClient.get(this.getUrl(url, params))
      .catch((error: any) => Observable.throw(error || 'Server error'));
  }

  private post(url, params = {}, options: any = {}) {
    return this.httpClient.post(this.getUrl(url), this.getQueryString(params))
      .catch((error: any) => Observable.throw(error || 'Server error'));
  }

  private getQueryString(parameters = null) {
    if (!parameters) {
      return '';
    }

    return Object.keys(parameters).reduce((array,key) => {
      array.push(key + '=' + encodeURIComponent(parameters[key]));
      return array;
    }, []).join('&');
  }

  private getUrl(url, options = null) {
    return this.url + url + '?' + this.getQueryString(options);
  }
}
