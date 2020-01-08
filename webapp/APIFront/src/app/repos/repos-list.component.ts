import { Component, OnInit  } from '@angular/core';
import { IResponse, IItem } from './repo';
import { RepositoriesService } from './repo.service';

@Component({
    templateUrl: './repos-list.component.html',
    styleUrls: ['./repos-list.component.css']
})
export class ReposListComponent {
    pageTitle: string = 'Repositories Search';

    _listFilter: string;
    errorMessage: any;
    get listFilter(): string {
      return this._listFilter;
    }
    set listFilter(value:string) {
      this._listFilter = value;
    }

    response: IResponse;
    items: IItem[] = [];

      constructor(private repositoriesService: RepositoriesService) {
        
      }
    
      OnSearch(): void {
        console.log(this.listFilter)
          this.repositoriesService.getRepos(this.listFilter).subscribe({
            next: response => {
              this.response = response
              this.items = response.items
              //console.log(response.total_count)
              console.log(this.items)
            },
            error: err => this.errorMessage = err
          });
          
  }
}