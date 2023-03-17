import { Component, OnInit, ElementRef, ViewChild } from '@angular/core';
import { Receipt, ReceiptDTO } from 'src/app/entities/receipt';
import { ReceiptService } from 'src/app/services/receipt.service';
import { Emitters } from 'src/app/emitters/emitters';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-qa',
  templateUrl: './qa.component.html',
  styleUrls: ['./qa.component.css']
})
export class QaComponent implements OnInit {
  @ViewChild("modelYear", {  }) modelYear: ElementRef;

  getLatestUnverifiedReceiptURL = "http://localhost:8080/api/receipts/unverified";
  getUserRolesUrl = "http://localhost:8080/api/user/roles";
  roles: string[] = [];

  receipt: Receipt = {
                      id: 0, 
                      model_year: 0, 
                      make: "", 
                      vin: "", 
                      first_name: "", 
                      last_name: "", 
                      state: "", 
                      date_added: "", 
                      qa_score: 0, 
                      qa_date: ""
                    };

  receiptDTO: ReceiptDTO = {
                      id: 0, 
                      model_year: 0, 
                      make: "", 
                      vin: "", 
                      first_name: "", 
                      last_name: "", 
                      state: "", 
                      date_added: "",
                      qa_score: 0, 
                      qa_date: ""
  } 

  constructor(
    private receiptService: ReceiptService,
    private http: HttpClient,
    private router: Router
  ) {}

  ngOnInit(): void {
    console.log("init")
    this.http.get(this.getUserRolesUrl, {withCredentials: true}).subscribe(
      (res: any) => {
        console.log(res)
        if (res != null) {
          this.roles = res;
        }
        console.log("got roles", this.roles)
        if (this.roles.indexOf("qa") == -1) {
          this.router.navigate(['/']);
        } else {
          this.refreshData();
          Emitters.inputEmitter.subscribe(
            () => {
              this.modelYear.nativeElement.focus()
            }
          );
        }
        
      },
      err => {
        console.error(err);
        console.log("did not get roles");
      }
    );
    
    
  }


  onSubmitTemplateBased(receiptFromForm: Receipt) { 
    this.receiptDTO.id = this.receipt.id
    this.receiptDTO.model_year = receiptFromForm.model_year
    this.receiptDTO.make = receiptFromForm.make
    this.receiptDTO.vin = receiptFromForm.vin
    this.receiptDTO.first_name = receiptFromForm.first_name
    this.receiptDTO.last_name = receiptFromForm.last_name
    this.receiptDTO.state = receiptFromForm.state
    this.receiptDTO.date_added = receiptFromForm.date_added
    this.receiptDTO.qa_score = receiptFromForm.qa_score
    this.receiptDTO.qa_date
    this.receiptService.updateVerifiedReceipt(this.receiptDTO)
      .subscribe(res => {
        console.log("receipt updated successfully: ", res);
      });
    Emitters.inputEmitter.emit();
  }

  updateReceipt() {
    console.log("receipt update started")
  }

  refreshData() {
    this.http.get(this.getLatestUnverifiedReceiptURL, {withCredentials: true}).subscribe(
      (res: any) => {
        this.receipt.id = res.id;
        this.receipt.model_year = res.model_year;
        this.receipt.make = res.make;
        this.receipt.vin = res.vin;
        this.receipt.first_name = res.first_name;
        this.receipt.last_name = res.last_name;
        this.receipt.state = res.state;
        this.receipt.date_added = res.date_added;
        this.receipt.qa_score = res.qa_score;
        this.receipt.qa_date = res.qa_date;
        this.modelYear.nativeElement.focus()
      },
      err => {
        console.error(err);        
      }
    );
  }
}
