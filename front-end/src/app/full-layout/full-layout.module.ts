import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FullLayoutComponent } from './full-layout/full-layout.component';
import { RouterModule, Routes } from '@angular/router';
import { ButtonsModule } from '@progress/kendo-angular-buttons';
import { LayoutModule, PanelBarModule } from '@progress/kendo-angular-layout';
import { InputsModule } from '@progress/kendo-angular-inputs';
import { FloatingLabelModule } from '@progress/kendo-angular-label';
import { DropDownsModule } from '@progress/kendo-angular-dropdowns';

const routes: Routes = [
  {
    path: '',
    component: FullLayoutComponent,
    children: [
      {
        path: '',
        redirectTo: '/home',
        pathMatch: 'full',
      },
      // {
      //   path: '',
      //   loadChildren: () =>
      //     import('./full-layout/dashboard/customers/customers.module').then(
      //       (m) => m.CustomersModule
      //     ),
      // },
    ],
  },
];
@NgModule({
  declarations: [FullLayoutComponent],
  imports: [
    CommonModule,
    RouterModule.forChild(routes),
    ButtonsModule,
    LayoutModule,
    InputsModule,
    FloatingLabelModule,
    PanelBarModule,
    DropDownsModule
  ],
})
export class FullLayoutModule {}
