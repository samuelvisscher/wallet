import { Component, Input, OnInit } from '@angular/core';
import { Wallet } from '../../app.datatypes';

@Component({
  selector: 'app-wallet',
  templateUrl: './wallet.component.html',
  styleUrls: ['./wallet.component.scss']
})
export class WalletComponent implements OnInit {

  @Input() wallet: Wallet;

  constructor() { }

  ngOnInit() {
  }

}
