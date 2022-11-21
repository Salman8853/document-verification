import { APP_INITIALIZER, NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { GridModule } from '@progress/kendo-angular-grid';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ButtonsModule } from '@progress/kendo-angular-buttons';
import { LayoutModule } from '@progress/kendo-angular-layout';
import { LabelModule } from '@progress/kendo-angular-label';
import { DropDownsModule } from '@progress/kendo-angular-dropdowns';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { InputsModule } from '@progress/kendo-angular-inputs';
import { AppConfig } from './services/app-config.service';
import { NotificationModule } from '@progress/kendo-angular-notification';
import { AuthGuard } from './services/auth/auth.service';
import { DialogsModule } from '@progress/kendo-angular-dialog';
import { LoaderComponent } from './full-layout/loader/loader.component';
import { IconsModule } from '@progress/kendo-angular-icons';
import { NavigationModule } from '@progress/kendo-angular-navigation';
import { TreeListModule } from '@progress/kendo-angular-treelist';
import { LoaderInterceptor } from './services/Interceptors/loader-interceptor.service';

function initConfig(config: AppConfig) {
  return () => config.ensureInit();
}

@NgModule({
  declarations: [AppComponent, LoaderComponent],
  imports: [
    BrowserModule,
    AppRoutingModule,
    GridModule,
    BrowserAnimationsModule,
    ButtonsModule,
    LayoutModule,
    LabelModule,
    DropDownsModule,
    HttpClientModule,
    InputsModule,
    NotificationModule,
    DialogsModule,
    IconsModule,
    NavigationModule,
    TreeListModule,
  ],
  providers: [
    AuthGuard,
    {
      provide: HTTP_INTERCEPTORS,
      useClass: LoaderInterceptor,
      multi: true,
    },
    AppConfig,
    {
      provide: APP_INITIALIZER,
      useFactory: initConfig,
      deps: [AppConfig],
      multi: true,
    },
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
