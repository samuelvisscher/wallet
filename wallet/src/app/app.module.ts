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
import { MatButtonModule } from '@angular/material/button';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatDialogModule } from '@angular/material/dialog';
import { RouterModule } from '@angular/router';
import { AppRoutes } from './app.routes';
import { ApiService, AppService, WalletService } from './services';
import { HttpClientModule } from '@angular/common/http';
import './rxjs-operators';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ReactiveFormsModule } from '@angular/forms';

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
  entryComponents: [
    CreateWalletComponent,
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    HttpClientModule,
    MatButtonModule,
    MatCheckboxModule,
    MatDialogModule,
    ReactiveFormsModule,
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
