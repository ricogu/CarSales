import { Injectable } from '@angular/core';
import {HttpClient, HttpParams} from "@angular/common/http";
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';
import {environment} from "../../environments/environment";
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})

export class OrderService {
  apiPort = environment['API_REST_PORT'];


  constructor(private http: HttpClient, private router: Router) {
  }

  getBatteries() {
    return this.http.get("/Batteries")
  }

  getWheels(batteryId : string ) {
    let params = new HttpParams();
    params = params.append('batteryId', batteryId);
    return this.http.get("/Wheels", {params : params})
  }

  getTires(wheelId : string ) {
    let params = new HttpParams();
    params = params.append('wheelId', wheelId);
    return this.http.get("/Tires", {params : params})
  }

  submitOrder(order : any) {
    return this.http.post("/Orders",order)
  }
}
