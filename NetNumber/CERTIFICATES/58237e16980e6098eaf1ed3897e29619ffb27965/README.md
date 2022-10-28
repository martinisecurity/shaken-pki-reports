# STIR/SHAKEN CA Ecosystem Compliance
## NetNumber

### Certificate 58237e16980e6098eaf1ed3897e29619ffb27965
Tested At: 2022-10-28 18:54:36 +0000 UTC\
Initial Validity Period: 31 day(s)\
Remaining Validity Period: 29 day(s)\
Subject: O=Baltimore-Washington Telephone Company, C=US, CN=Baltimore-Washington Telephone Company SHAKEN cert 8697\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1

Link: https://bwt.io/8697a2.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIICvTCCAkOgAwIBAgIJALhWb4Z4c%2FMRMAoGCCqGSM49BAMCMIGOMTAwLgYDVQQDDCdOZXROdW1iZXIgU0hBS0VOIFJvb3QgSW50ZXJtZWRpYXRlIENBIDExCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1OZXROdW1iZXIgSW5jMQswCQYDVQQLDAJVUzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDAeFw0yMjEwMjYxNTAwMDBaFw0yMjExMjYxNDU5NTlaMIGAMUAwPgYDVQQDDDdCYWx0aW1vcmUtV2FzaGluZ3RvbiBUZWxlcGhvbmUgQ29tcGFueSBTSEFLRU4gY2VydCA4Njk3MQswCQYDVQQGEwJVUzEvMC0GA1UECgwmQmFsdGltb3JlLVdhc2hpbmd0b24gVGVsZXBob25lIENvbXBhbnkwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAR7PqgstTJ0GphOuUb64MPJkWWX6Nd%2FNz9kYzpj%2BjtDTPIPk2EMAzGCSNNBFuRF0obuipQkJf6mYYHRZPSkS%2BqBo4GVMIGSMBYGCCsGAQUFBwEaBAowCKAGFgQ4Njk3MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB8GA1UdIwQYMBaAFHEvyILcqOAi49%2BgaPn4XlyE3uW9MB0GA1UdDgQWBBS8CFGQU1KbdE7xcbVIUHRr4NPURTAaBgNVHSABAf8EEDAOMAwGCmCGSAGG%2FwkBAQEwCgYIKoZIzj0EAwIDaAAwZQIxAPMy7RpVEhH18GHeIZbV%2BQvu%2FRDW7nbpd%2BOCwX5h9CdEWIA0eXEUG5DMIRQ3PW8wDQIwaKCpuLFr0PrWi%2BWSXfyBaq%2FS9eTU6HQYe%2B76MfdBb1%2BuxtDQLdioT1ChSrerABEJ)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| n_sti_certificate_policy_critical | notice | ATIS-1000080 | STI certificates should contain a CertificatePolicies extension marked uncritical |
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 8697' |
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_crl_distribution | error | ATIS-1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |

Generated: 28/10/2022 at 18:55:01