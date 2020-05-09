import { Component, OnInit } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {ReportService} from "./report.service";

@Component({
  selector: 'app-report',
  templateUrl: './report.component.html',
  styleUrls: ['./report.component.css']
})
export class ReportComponent implements OnInit {

  orders;

  reportSuccess;

  constructor(private reportService : ReportService) { }

  ngOnInit(): void {
    this.reportService.getOrders().subscribe(
      data => {this.orders = data,this.reportSuccess = true},
      error => {console.error('There was an error!', error), this.reportSuccess= false}
    )
  }

}
