import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { OrderComponent } from './order/order.component';
import { HttpClientModule } from "@angular/common/http";
import { NotifierModule } from "angular-notifier";
import {FormsModule} from "@angular/forms";

@NgModule({
  declarations: [
    AppComponent,
    OrderComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    NotifierModule.withConfig(
      {
        position: {
          horizontal: {
            position: 'middle',
            distance: 12

          },

          vertical: {
            position: 'top',
            distance: 12,
            gap: 10

          }

        },
        theme: "material",
        behaviour: {
          autoHide: 3000
        }
      }
    ),
    FormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})

export class AppModule {
  title = "Car Ordering System"
}
