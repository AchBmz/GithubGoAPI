import { NgModule } from '@angular/core';
import { ReposListComponent } from './repos-list.component';
import { RouterModule } from '@angular/router';
import { SharedModule } from '../shared/shared.module';

@NgModule({
  declarations: [
    ReposListComponent
  ],
  imports: [
    RouterModule.forChild([
      { path: 'repos', component: ReposListComponent },
    ]),
    SharedModule
  ]
})
export class RepoModule { }
