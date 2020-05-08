# table and data for batteries
drop table if exists `Batteries`;
create table Batteries
(
	Battery varchar(64) not null comment 'Name of Battery'
		primary key,
	Price int default 0 not null comment 'Additional Price of Battery Model '
);

insert into Batteries (Battery, Price) values ('40 kwh', 0);
insert into Batteries (Battery, Price) values ('60 kwh', 2500);
insert into Batteries (Battery, Price) values ('80 kwh', 6000);

# table and data for tires
drop table if exists `Tires`;
create table Tires
(
	Tire varchar(32) not null comment 'Name of the Tire profile'
		primary key,
	Price int not null comment 'Additional Price of the tire'
);

insert into Tires (Tire, Price) values ('Eco', 0);
insert into Tires (Tire, Price) values ('Performance', 80);
insert into Tires (Tire, Price) values ('Racing', 150);

# table and data for wheels
drop table if exists `Wheels`;
create table Wheels
(
	Wheel varchar(32) not null comment 'Name of the Wheel'
		primary key,
	Price int not null comment 'Additional Price of the wheel model'
);

insert into Wheels (Wheel, Price) values ('Model 1', 0);
insert into Wheels (Wheel, Price) values ('Model 2', 150);
insert into Wheels (Wheel, Price) values ('Model 3', 350);

# table and data for TiresAvailability, relationship between tire and wheel
drop table if exists `TiresAvailability`;
create table TiresAvailability
(
	TireName varchar(32) not null,
	WheelName varchar(32) not null,
	constraint TiresAvailability_Tires_Tire_fk
		foreign key (TireName) references Tires (Tire),
	constraint TiresAvailability_Wheels_Wheel_fk
		foreign key (WheelName) references Wheels (Wheel)
)
comment 'Availability of tires depends on the wheel model choice';

insert into TiresAvailability (TireName, WheelName) values ('Eco', 'model 1');
insert into TiresAvailability (TireName, WheelName) values ('Eco', 'model 2');
insert into TiresAvailability (TireName, WheelName) values ('Eco', 'model 3');
insert into TiresAvailability (TireName, WheelName) values ('Performance', 'model 2');
insert into TiresAvailability (TireName, WheelName) values ('Performance', 'model 3');
insert into TiresAvailability (TireName, WheelName) values ('Racing', 'model 3');

# table and data for TiresAvailability, relationship between battery and wheel
drop table if exists `WheelAvailability`;
create table WheelAvailability
(
	BatteryName varchar(64) not null,
	WheelName varchar(32) not null,
	constraint WheelAvaibility_Batteries_Battery_fk
		foreign key (BatteryName) references Batteries (Battery),
	constraint WheelAvaibility_Wheels_Wheel_fk
		foreign key (WheelName) references Wheels (Wheel)
)
comment 'Availability of wheel depends on the battery choice';

insert into WheelAvailability (BatteryName, WheelName) values ('40 kwh', 'model 1');
insert into WheelAvailability (BatteryName, WheelName) values ('40 kwh', 'model 2');
insert into WheelAvailability (BatteryName, WheelName) values ('60 kwh', 'model 1');
insert into WheelAvailability (BatteryName, WheelName) values ('60 kwh', 'model 2');
insert into WheelAvailability (BatteryName, WheelName) values ('60 kwh', 'model 3');
insert into WheelAvailability (BatteryName, WheelName) values ('80 kwh', 'model 1');
insert into WheelAvailability (BatteryName, WheelName) values ('80 kwh', 'model 2');
insert into WheelAvailability (BatteryName, WheelName) values ('80 kwh', 'model 3');
