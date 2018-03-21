import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { WalletService } from '../../services';
import { Wallet } from '../../app.datatypes';
import { MatDialog } from '@angular/material';
import { CreateWalletComponent } from '../create-wallet/create-wallet.component';

@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.scss']
})
export class SidebarComponent implements OnInit {

  @Output() onSelect = new EventEmitter();

  wallets: Wallet[];

  constructor(
    public dialog: MatDialog,
    private walletService: WalletService,
  ) { }

  ngOnInit() {
    this.walletService.wallets.subscribe(wallets => {
      this.wallets = wallets;
      if (this.wallets && this.wallets[0]) {
        this.open(this.wallets[0]);
      }
    });
  }

  createWallet() {
    this.dialog.open(CreateWalletComponent, { width: '700px' });
  }

  open(wallet: Wallet) {
    this.onSelect.emit(wallet);
  }
}
