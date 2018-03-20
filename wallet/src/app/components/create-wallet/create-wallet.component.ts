import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-create-wallet',
  templateUrl: './create-wallet.component.html',
  styleUrls: ['./create-wallet.component.css']
})
export class CreateWalletComponent implements OnInit {

  step = 1;

  constructor() { }

  ngOnInit() {
  }

  handleStepOne(values: any) {
    this.step = 2;
  }

  handleStepTwo() {
    this.step = 3;
  }
}
