import { Component } from '@angular/core';
import { Wallet } from './app.datatypes';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {

  wallet: Wallet;

  handleOnSelect(wallet: Wallet) {
    this.wallet = wallet;
  }
}
