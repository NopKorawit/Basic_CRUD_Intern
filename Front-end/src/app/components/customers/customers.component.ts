import { Component, OnInit } from '@angular/core';
import { Customer } from 'src/app/models/customer';
import { CUSTOMERS } from 'src/app/data/mock-customer';
import { map, Observable, startWith } from 'rxjs';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { NgbInputDatepickerConfig, NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { CustomerService } from '../../service/customer.service';

@Component({
  selector: 'app-customers',
  templateUrl: './customers.component.html',
  styleUrls: ['./customers.component.css'],
  providers: [NgbInputDatepickerConfig],
})
export class CustomersComponent implements OnInit {
  customers = CUSTOMERS;
  errorMessage = '';
  errorClass = '';

  getDateString(birthday: Date) {
    const result = new Date(birthday).toLocaleDateString('th');
    return result;
  }
  calString(dateOfBirth: string) {
    const now = Date.now();
    const dob = new Date(dateOfBirth).getTime();
    const timeDiff = Math.abs(now - dob);
    const result = Math.floor(timeDiff / (1000 * 3600 * 24) / 365);

    return result;
  }
  calculateAge(dob: number) {
    const now = Date.now();
    const timeDiff = Math.abs(now - dob);
    const result = Math.floor(timeDiff / (1000 * 3600 * 24) / 365);

    return result;
  }
  customerList: any;
  getCustomers() {
    this.service.getCustomer().subscribe((result) => {
      if (result.message === 'successful') {
        this.customerList = result.data;
      }
    });
  }
  getCustomer(id?: number | string) {
    const result = this.service.getCustomer(id);
    return result;
  }
  closeResult = '';
  constructor(
    private modalService: NgbModal,
    private service: CustomerService,

    config: NgbInputDatepickerConfig
  ) {
    const now = new Date();
    config.minDate = { year: 1980, month: 1, day: 1 };
    config.maxDate = {
      year: now.getFullYear(),
      month: now.getMonth(),
      day: now.getDay(),
    };
  }
  open(content: any, id?: number) {
    if (id) {
      console.log(id);
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
    id: new FormControl(1),
    firstName: new FormControl(
      '',
      Validators.compose([Validators.required, Validators.minLength(4)])
    ),
    lastName: new FormControl(
      '',
      Validators.compose([Validators.required, Validators.minLength(4)])
    ),
    birthday: new FormControl('', Validators.required),
    address: new FormControl('', Validators.required),
    //  age: new FormControl()
  });
  saveResponse: any;
  tempData: Customer = {
    firstName: '',
    lastName: '',
    birthday: '',
    address: '',
  };
  saveCustomer() {
    if (this.customerForm.valid) {
      this.tempData = this.customerForm.getRawValue();
      this.tempData.birthday = new Date(
        `${this.customerForm.get('birthday')?.value.year}-${
          this.customerForm.get('birthday')?.value.month
        }-${this.customerForm.get('birthday')?.value.day}`
      ).toISOString();

      this.service.createCustomer(this.tempData).subscribe((result) => {
        this.saveResponse = result;
        if (this.saveResponse.result == 'pass') {
          this.errorMessage = 'Saved';
          this.errorClass = 'success';
          this.modalService.dismissAll();
        } else {
          this.errorMessage = 'Failed to save';
        }
      });
      console.log(this.tempData);
    }
  }

  // get age() {
  //   return this.calculateAge(this.customerForm.getRawValue().birthday);
  // }
  get age() {
    return this.calculateAge(
      new Date(
        `${this.customerForm.get('birthday')?.value.year}-${
          this.customerForm.get('birthday')?.value.month
        }-${this.customerForm.get('birthday')?.value.day}`
      ).getTime()
    );
  }
  editData: any;
  loadEditData(id: number) {
    this.service.getCustomer(id).subscribe((result) => {
      this.editData = result;
    });
  }
  functionEdit(id: number) {
    this.loadEditData(id);
  }
}
