import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError, tap, map } from 'rxjs/operators';

import { IResponse, IItem } from './repo';

@Injectable({
  providedIn: 'root'
})
export class RepositoriesService {
  private searchByLanguageUrl = 'http://localhost:3000/?language=';
  constructor(private http: HttpClient) { }

  getRepos(language: string): Observable<IResponse> {
    console.log('==> ' + language)
    return this.http.get<IResponse>(this.searchByLanguageUrl+language)
      .pipe(
        tap(data => console.log('JSON: ' + JSON.stringify(data))),
        catchError(this.handleError)
      );
  }

  private handleError(err: HttpErrorResponse) {
    let errorMessage = '';
    if (err.error instanceof ErrorEvent) {
      errorMessage = `An error occurred: ${err.error.message}`;
    } else {
      errorMessage = `Server returned code: ${err.status}, error message is: ${err.message}`;
    }
    console.error(errorMessage);
    return throwError(errorMessage);
  }

}
