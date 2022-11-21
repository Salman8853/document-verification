import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from './services/auth/auth.service';

const routes: Routes = [
  {
    path: '',
    redirectTo: '/login',
    pathMatch: 'full'
  },
  {
    path: 'login',
    loadChildren: () =>
      import('./full-layout/login/login.module').then((m) => m.LoginModule),
  },
  // {
  //   path: 'home',
  //   canLoad: [],
  //   loadChildren: () =>
  //     import('./full-layout/full-layout.module').then(
  //       (m) => m.FullLayoutModule
  //     ),
  // },
  {
    path: 'home',
    canLoad: [AuthGuard],
    loadChildren: () =>
      import('./full-layout/full-layout.module').then(
        (m) => m.FullLayoutModule
      ),
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
