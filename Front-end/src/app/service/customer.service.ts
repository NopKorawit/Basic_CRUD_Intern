import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Customer } from '../models/customer';
import { HttpResult } from '../models/http';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root',
})
export class CustomerService {
  subscribeInstant: any;
  constructor(private http: HttpClient) {}

  getCustomer(Id?: string | number): Observable<HttpResult<Customer[]>> {
    const url = Id
      ? `${environment.baseApi}/customer/${Id}`
      : `${environment.baseApi}/customer`;
    this.subscribeInstant = this.http.get(url);
    return this.subscribeInstant;
  }

  deleteCustomer(id: number): Observable<HttpResult<Customer[]>> {
    this.subscribeInstant = this.http.delete(
      `${environment.baseApi}/customer/${id}`
    );
    return this.subscribeInstant;
  }
  createCustomer(payload: Customer): Observable<HttpResult<string>> {
    this.subscribeInstant = this.http.post(
      `${environment.baseApi}/customer`,
      payload
    );
    return this.subscribeInstant;
  }
}
