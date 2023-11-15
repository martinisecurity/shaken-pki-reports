# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Number Access LLC SHAKEN 343J

Tested At: 15 Nov 23 16:12 UTC\
Initial Validity Period: 42 day(s)\
Remaining Validity Period: -184 day(s)\
Subject: L=Bettendorf, ST=Iowa, O=Number Access LLC, C=US, CN=Number Access LLC SHAKEN 343J\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://ss-public-certs.numberaccess.net/na-20230403131109-20230515131109.pem

[View certificate details](https://understandingwebpki.com/?cert=MIICsjCCAjegAwIBAgIJAODaXt7GZtzPMAoGCCqGSM49BAMCMIGOMTAwLgYDVQQDDCdOZXROdW1iZXIgU0hBS0VOIFJvb3QgSW50ZXJtZWRpYXRlIENBIDExCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1OZXROdW1iZXIgSW5jMQswCQYDVQQLDAJVUzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDAeFw0yMzA0MDMxMzExMDlaFw0yMzA1MTUxMzExMDlaMHUxJjAkBgNVBAMMHU51bWJlciBBY2Nlc3MgTExDIFNIQUtFTiAzNDNKMQswCQYDVQQGEwJVUzEaMBgGA1UECgwRTnVtYmVyIEFjY2VzcyBMTEMxDTALBgNVBAgMBElvd2ExEzARBgNVBAcMCkJldHRlbmRvcmYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATpL7xgZNdK%2FlaGTSWwLyiuhVCxCIZCM4T%2BuLIDuQae%2BuP2BOu6qu%2FI3arWpS1byYPLNQSruM4%2Bd8TIPOkygTDlo4GVMIGSMBYGCCsGAQUFBwEaBAowCKAGFgQzNDNKMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB8GA1UdIwQYMBaAFHEvyILcqOAi49%2BgaPn4XlyE3uW9MB0GA1UdDgQWBBRKOwGqCRp9XLRXQ37%2B3z7QlwbyjTAaBgNVHSABAf8EEDAOMAwGCmCGSAGG%2FwkBAQEwCgYIKoZIzj0EAwIDaQAwZgIxAMA9c14vkn7NjH5vLzX5gKfCqDadwBZbgIBWzQv7fxjCfO1Z8HgbZosmW5wlVKH0SQIxAPb5FFu5saBgTG01Af4jK61bvSvRhxsOgjtHJ07KeiBqUod2LIRF%2FvzvGIl%2FDOK3xA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [n_atis_certificate_policy_critical](../../ISSUES/n_atis_certificate_policy_critical/README.md) | notice | ATIS1000080 | STI certificates should contain a CertificatePolicies extension marked uncritical |
| [e_atis_crl_distribution](../../ISSUES/e_atis_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |


Generated: 15 Nov 23 16:51 UTC