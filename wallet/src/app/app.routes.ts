import { WalletComponent } from './components/wallet/wallet.component';


export const AppRoutes = [
  {
    path: '',
    component: WalletComponent,
  },
  {
    path: 'wallet/:id',
    component: WalletComponent,
  },
];
