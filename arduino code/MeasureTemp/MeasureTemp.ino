#include <SPI.h>
#include "Adafruit_MAX31855.h"

// Default connection is using software SPI, but comment and uncomment one of
// the two examples below to switch between software SPI and hardware SPI:

// Example creating a thermocouple instance with software SPI on any three
// digital IO pins.
#define MAXDO   2
#define MAXCS   3
#define MAXCLK  4


#define MAXDO2   11
#define MAXCS2   12
#define MAXCLK2  13

// initialize the Thermocouples
Adafruit_MAX31855 thermocouple(MAXCLK, MAXCS, MAXDO);
Adafruit_MAX31855 thermocouple2(MAXCLK2, MAXCS2, MAXDO2);

// Example creating a thermocouple instance with hardware SPI (Uno/Mega only)
// on a given CS pin.
//#define MAXCS   10
//Adafruit_MAX31855 thermocouple(MAXCS);

void setup() {  
  Serial.begin(9600);
  
  Serial.println("MAX31855 test. Hi! I'm an Arduino");
  // wait for MAX chip to stabilize
  delay(500);
}

void loop() {
  // basic readout test, just print the current temp
   Serial.print("Internal Temp = ");
   Serial.println(thermocouple.readInternal());
   
   Serial.print("Internal Temp 2 = ");
   Serial.println(thermocouple2.readInternal());
   
   double c = thermocouple.readCelsius();
   if (isnan(c)) {
     Serial.println("Something wrong with thermocouple!");
   } else {
     Serial.print("C = "); 
     Serial.println(c);
   }
   
   double c2 = thermocouple2.readCelsius();
   if (isnan(c2)) {
     Serial.println("Something wrong with thermocouple2!");
   } else {
     Serial.print("C2 = "); 
     Serial.println(c2);
   }
   
   //Serial.print("F = ");
   //Serial.println(thermocouple.readFarenheit());
 
   delay(1000);
}
