#drop exisitng tables
drop table if exists `TiresAvailability`;
drop table if exists `WheelAvailability`;
drop table if exists `Orders`;
drop table if exists `Batteries`;
drop table if exists `Tires`;
drop table if exists `Wheels`;


# table and data for batteries
create table Batteries
(
    ID int auto_increment comment 'ID of Battery',
	Battery varchar(64) not null comment 'Name of Battery',
	Price int default 0 not null comment 'Additional Price of Battery Model',
    constraint Batteries._pk
    primary key (ID)
);

insert into Batteries (Battery, Price) values ('40 kwh', 0);
insert into Batteries (Battery, Price) values ('60 kwh', 2500);
insert into Batteries (Battery, Price) values ('80 kwh', 6000);

# table and data for tires
create table Tires
(
    ID int auto_increment comment 'ID of Tires',
	Tire varchar(32) not null comment 'Name of the Tire profile',
	Price int not null comment 'Additional Price of the tire',
    constraint Tires._pk
    primary key (ID)
);

insert into Tires (Tire, Price) values ('Eco', 0);
insert into Tires (Tire, Price) values ('Performance', 80);
insert into Tires (Tire, Price) values ('Racing', 150);

# table and data for wheels
create table Wheels
(
    ID int auto_increment comment 'ID of Wheels',
	Wheel varchar(32) not null comment 'Name of the Wheel',
	Price int not null comment 'Additional Price of the wheel model',
    constraint Wheels._pk
    primary key (ID)
);

insert into Wheels (Wheel, Price) values ('Model 1', 0);
insert into Wheels (Wheel, Price) values ('Model 2', 150);
insert into Wheels (Wheel, Price) values ('Model 3', 350);

# table and data for TiresAvailability, relationship between tire and wheel
create table TiresAvailability
(
	TireID int not null ,
	WheelID int not null,
	constraint TiresAvailability_Tires_Tire_fk
		foreign key (TireID) references Tires (ID),
	constraint TiresAvailability_Wheels_Wheel_fk
		foreign key (WheelID) references Wheels (ID)
)
comment 'Availability of tires depends on the wheel model choice';

insert into TiresAvailability (TireID, WheelID) values (1, 1);
insert into TiresAvailability (TireID, WheelID) values (1, 2);
insert into TiresAvailability (TireID, WheelID) values (1, 3);
insert into TiresAvailability (TireID, WheelID) values (2, 2);
insert into TiresAvailability (TireID, WheelID) values (2, 3);
insert into TiresAvailability (TireID, WheelID) values (3, 3);

# table and data for TiresAvailability, relationship between battery and wheel
create table WheelAvailability
(
	BatteryID int not null,
	WheelID int not null,
	constraint WheelAvaibility_Batteries_Battery_fk
		foreign key (BatteryID) references Batteries (ID),
	constraint WheelAvaibility_Wheels_Wheel_fk
		foreign key (WheelID) references Wheels (ID)
)
comment 'Availability of wheel depends on the battery choice';

insert into WheelAvailability (BatteryID, WheelID) values (1, 1);
insert into WheelAvailability (BatteryID, WheelID) values (1, 2);
insert into WheelAvailability (BatteryID, WheelID) values (2, 1);
insert into WheelAvailability (BatteryID, WheelID) values (2, 2);
insert into WheelAvailability (BatteryID, WheelID) values (2, 3);
insert into WheelAvailability (BatteryID, WheelID) values (3, 1);
insert into WheelAvailability (BatteryID, WheelID) values (3, 2);
insert into WheelAvailability (BatteryID, WheelID) values (3, 3);

#Tables recording car orders
create table Orders
(
    ID int auto_increment,
    TireID int not null,
    WheelID int not null,
    BatteryID int not null,
    Discount bool default false not null,
    NetCost int not null,
    FinalCost int not null,
    constraint Orders_pk
        primary key (ID),
    constraint Orders_Batteries_ID_fk
        foreign key (BatteryID) references Batteries (ID),
    constraint Orders_Tires_ID_fk
        foreign key (TireID) references Tires (ID),
    constraint Orders_Wheels_ID_fk
        foreign key (WheelID) references Wheels (ID)
);
