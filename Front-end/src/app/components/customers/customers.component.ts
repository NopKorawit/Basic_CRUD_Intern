import { Component, OnInit, PipeTransform } from '@angular/core';
import { Customer } from 'src/app/models/customer';
import { CUSTOMERS } from 'src/app/data/mock-customer';
import { map, Observable, startWith } from 'rxjs';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { NgbDateAdapter, NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { CustomerService } from '../../service/customer.service';

@Component({
  selector: 'app-customers',
  templateUrl: './customers.component.html',
  styleUrls: ['./customers.component.css'],
})
export class CustomersComponent implements OnInit {
  customers = CUSTOMERS;
  errorMessage = '';
  errorClass = '';

  getDateString(dateOfBirth: Date) {
    const result = new Date(dateOfBirth).toLocaleDateString('th');
    return result;
  }

  calculateAge(dateOfBirth: Date) {
    const dob = dateOfBirth.getTime();
    const now = Date.now();
    const timeDiff = Math.abs(now - dob);
    const result = Math.floor(timeDiff / (1000 * 3600 * 24) / 365);

    return result;
  }
  customerList: any;
  getCustomers() {
    this.service.getCustomer().subscribe((result) => {
      this.customerList = result;
    });
  }
  getCustomer(id: number) {
    const result: Customer = this.customers[id];
    return result;
  }
  closeResult = '';
  constructor(
    private modalService: NgbModal,
    private service: CustomerService,
    private dateAdapter: NgbDateAdapter<string>
  ) {}
  open(content: any, customer?: Customer) {
    if (customer) {
      console.log(customer.id);
    }
    this.modalService
      .open(content, {
        ariaLabelledBy: 'modal-basic-title',
        centered: true,
        animation: true,
      })
      .result.then((result) => {
        this.closeResult = `Closed with: ${result}`;
      });
  }

  ngOnInit(): void {
    this.getCustomers();
  }
  customerForm = new FormGroup({
    id: new FormControl({ value: 1 }),
    firstName: new FormControl(
      '',
      Validators.compose([Validators.required, Validators.minLength(4)])
    ),
    lastName: new FormControl(
      '',
      Validators.compose([Validators.required, Validators.minLength(4)])
    ),
    dateOfBirth: new FormControl('', Validators.required),
    address: new FormControl('', Validators.required),
    //  age: new FormControl()
  });
  saveResponse: any;
  saveCustomer() {
    console.log(this.customerForm.getRawValue());
    // console.log(
    //   'Age :' +
    //     this.calculateAge(this.customerForm.getRawValue().dateOfBirth) +
    //     'ปี'
    // );

    if (this.customerForm.valid) {
      this.service
        .createCustomer(this.customerForm.getRawValue())
        .subscribe((result) => {
          this.saveResponse = result;
          if (this.saveResponse.result == 'pass') {
            this.errorMessage = 'Saved';
            this.errorClass = 'success';
          } else {
            this.errorMessage = 'Failed to save';
          }
        });
    }
  }

  // get age() {
  //   return this.calculateAge(this.customerForm.getRawValue().dateOfBirth);
  // }
  loadEditData(id: any) {
    console.log(id);
  }
  functionEdit(id: any) {
    this.loadEditData(id);
  }
}
