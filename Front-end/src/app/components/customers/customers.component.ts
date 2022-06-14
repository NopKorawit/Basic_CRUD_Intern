import { Component, OnInit, PipeTransform } from '@angular/core';

import { Customer } from 'src/app/models/customer';
import { CUSTOMERS } from 'src/app/data/mock-customer';
import { map, Observable, startWith } from 'rxjs';
import { FormControl } from '@angular/forms';

import { NgbModal, ModalDismissReasons } from '@ng-bootstrap/ng-bootstrap';

@Component({
  selector: 'app-customers',
  templateUrl: './customers.component.html',
  styleUrls: ['./customers.component.css'],
})
export class CustomersComponent implements OnInit {
  customers = CUSTOMERS;

  getDateString(dateOfBirth: number) {
    const result = new Date(dateOfBirth).toLocaleDateString('th');
    return result;
  }

  calculateAge(dateOfBirth: number) {
    const timeDiff = Math.abs(Date.now() - dateOfBirth);
    const result = Math.floor(timeDiff / (1000 * 3600 * 24) / 365);

    return result;
  }
  getCustomer(id: number) {
    const result: Customer = this.customers[id];
    return result;
  }
  closeResult = '';
  constructor(private modalService: NgbModal) {}
  open(content: any, customer?: Customer) {
    this.modalService
      .open(content, { ariaLabelledBy: 'modal-basic-title', centered: true })
      .result.then(
        (result) => {
          this.closeResult = `Closed with: ${result}`;
        },
        (reason) => {
          this.closeResult = `Dismissed ${this.getDismissReason(reason)}`;
        }
      );
  }
  private getDismissReason(reason: any): string {
    if (reason === ModalDismissReasons.ESC) {
      return 'by pressing ESC';
    } else if (reason === ModalDismissReasons.BACKDROP_CLICK) {
      return 'by clicking on a backdrop';
    } else {
      return `with: ${reason}`;
    }
  }
  ngOnInit(): void {}
}
