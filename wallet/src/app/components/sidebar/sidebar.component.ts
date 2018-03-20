import { Component, OnInit } from '@angular/core';
import { WalletService } from '../../services';
import { MatDialog } from '@angular/material';
import { CreateWalletComponent } from '../create-wallet/create-wallet.component';

@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.scss']
})
export class SidebarComponent implements OnInit {

  constructor(
    public dialog: MatDialog,
  ) { }

  ngOnInit() {
  }

  createWallet() {
    this.dialog.open(CreateWalletComponent, { width: '700px' });
  }
}
