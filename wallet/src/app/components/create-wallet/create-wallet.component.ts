import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../services';

@Component({
  selector: 'app-create-wallet',
  templateUrl: './create-wallet.component.html',
})
export class CreateWalletComponent implements OnInit {

  label: string;
  seed: string;
  step = 1;

  constructor(
    private apiService: ApiService,
  ) { }

  ngOnInit() {

  }

  handleStepOne(values: any) {
    this.label = values.name;
    this.step = 2;
  }

  handleStepTwo() {
    this.apiService.getWalletsSeed().subscribe(
      seed => {
        if (seed) {
          this.seed = seed;
          this.step = 3;
        }
      }
    )
  }

  handleStepThree() {
    const request = {
      label: this.label,
      seed: this.seed,
      aCount: 1,
      encrypted: false,
      password: null,
    };

    this.apiService.postWalletsNew(request).subscribe(response => console.log(response));
  }
}
