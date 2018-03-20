import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { AppComponent } from './app.component';
import { SidebarComponent } from './components/sidebar/sidebar.component';
import { WalletComponent } from './components/wallet/wallet.component';
import { KittenTileComponent } from './components/wallet/kitten-tile/kitten-tile.component';
import { AddressPanelComponent } from './components/wallet/address-panel/address-panel.component';
import { CreateWalletComponent } from './components/create-wallet/create-wallet.component';
import { SetupComponent } from './components/create-wallet/setup/setup.component';
import { WarningComponent } from './components/create-wallet/warning/warning.component';
import { ShowSeedComponent } from './components/create-wallet/show-seed/show-seed.component';
import { FeedComponent } from './components/feed/feed.component';
import { BreedComponent } from './components/breed/breed.component';
import { KittenDetailComponent } from './components/kitten-detail/kitten-detail.component';
import { RouterModule } from '@angular/router';
import { AppRoutes } from './app.routes';
import { ApiService, AppService, WalletService } from './services';
import './rxjs-operators';

@NgModule({
  declarations: [
    AppComponent,
    SidebarComponent,
    WalletComponent,
    KittenTileComponent,
    AddressPanelComponent,
    CreateWalletComponent,
    SetupComponent,
    WarningComponent,
    ShowSeedComponent,
    FeedComponent,
    BreedComponent,
    KittenDetailComponent,
  ],
  imports: [
    BrowserModule,
    RouterModule.forRoot(AppRoutes),
  ],
  providers: [
    AppService,
    ApiService,
    WalletService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
