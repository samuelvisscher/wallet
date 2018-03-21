import { Injectable } from '@angular/core';
import { Subject } from 'rxjs/Subject';
import { Wallet } from '../app.datatypes';
import { BehaviorSubject } from 'rxjs/BehaviorSubject';
import { ApiService } from './api.service';
import { Observable } from 'rxjs/Observable';

@Injectable()
export class WalletService {

  private walletsSubject: Subject<Wallet[]> = new BehaviorSubject<Wallet[]>([]);

  get wallets(): Observable<Wallet[]> {
    return this.walletsSubject.asObservable();
  }

  constructor(
    private apiService: ApiService,
  ) {
    this.loadData();
  }

  loadData() {
    this.apiService.getWalletsList().subscribe(wallets => this.walletsSubject.next(wallets));
  }
}
