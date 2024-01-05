# SLAs - Service Level Agreements
Agreements made with the costumers about the reliability of your services
A SLA needs to have conquencences if it is violated/fails.
If it is a service, giving partial refunds as a way of compensation, or extra 
resources can be a good way of compensation


# SLOs - Service Level Objectives
Catch an issue before it affects your SLA, so you have enough time to fix it
SLOs are always stronger than SLAs
Once the SLOs are violated, you should remove risks from your service:
+ Slowing down the rate of change to the system;
+ Eliminating risks by doing fewer pushes;
+ Devoting engineering and automation efforts to reducing and eliminating areas of risks.

## How to measure reliability
  + Time to start playing:
    * Latency measurement 
 
  * No interruptions or isses with playback

# SLIs - Service Level Indicators
Requesting the latency is a quantitative measurement/metric 

## How to set SLOs for SLIs
A SLO is a target that we define, so it depends a lot on the parameter youy choose.

Ex: target SLO = 99% -> 99% of the requests will be served within 300ms

# Edge Cases...
