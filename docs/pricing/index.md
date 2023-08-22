#Pricing for Ingress

Pricing cost is atleast (per million requests at 10kb storage each):

Planetscale INSERTS: $1.5
Planetscale Reads: $1
Planetscale Storage: $25 (goes down to $1.5 in scaler pro)
Lambda: $1
S3: $5.24 (PUT)
S3 monthly: $0.23 (GET)


Ingress price: $1.5 + $1 + $5.24 = $7.74 (one time)
Ingress Query storage: $25/month
Cold Storage price: $1 + $0.23 = $1.23/month

