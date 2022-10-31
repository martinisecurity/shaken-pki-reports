# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Baltimore-Washington Telephone Company SHAKEN cert 8697

Tested At: 31 Oct 22 20:24 UTC\
Initial Validity Period: 31 day(s)\
Remaining Validity Period: 26 day(s)\
Subject: O=Baltimore-Washington Telephone Company, C=US, CN=Baltimore-Washington Telephone Company SHAKEN cert 8697\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://bwt.io/8697a2.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIICvTCCAkOgAwIBAgIJALhWb4Z4c%2FMRMAoGCCqGSM49BAMCMIGOMTAwLgYDVQQDDCdOZXROdW1iZXIgU0hBS0VOIFJvb3QgSW50ZXJtZWRpYXRlIENBIDExCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1OZXROdW1iZXIgSW5jMQswCQYDVQQLDAJVUzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDAeFw0yMjEwMjYxNTAwMDBaFw0yMjExMjYxNDU5NTlaMIGAMUAwPgYDVQQDDDdCYWx0aW1vcmUtV2FzaGluZ3RvbiBUZWxlcGhvbmUgQ29tcGFueSBTSEFLRU4gY2VydCA4Njk3MQswCQYDVQQGEwJVUzEvMC0GA1UECgwmQmFsdGltb3JlLVdhc2hpbmd0b24gVGVsZXBob25lIENvbXBhbnkwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAR7PqgstTJ0GphOuUb64MPJkWWX6Nd%2FNz9kYzpj%2BjtDTPIPk2EMAzGCSNNBFuRF0obuipQkJf6mYYHRZPSkS%2BqBo4GVMIGSMBYGCCsGAQUFBwEaBAowCKAGFgQ4Njk3MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB8GA1UdIwQYMBaAFHEvyILcqOAi49%2BgaPn4XlyE3uW9MB0GA1UdDgQWBBS8CFGQU1KbdE7xcbVIUHRr4NPURTAaBgNVHSABAf8EEDAOMAwGCmCGSAGG%2FwkBAQEwCgYIKoZIzj0EAwIDaAAwZQIxAPMy7RpVEhH18GHeIZbV%2BQvu%2FRDW7nbpd%2BOCwX5h9CdEWIA0eXEUG5DMIRQ3PW8wDQIwaKCpuLFr0PrWi%2BWSXfyBaq%2FS9eTU6HQYe%2B76MfdBb1%2BuxtDQLdioT1ChSrerABEJ)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_sti_crl_distribution](../../ISSUES/e_sti_crl_distribution/README.md) | error | ATIS&#x2011;1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [n_sti_certificate_policy_critical](../../ISSUES/n_sti_certificate_policy_critical/README.md) | notice | ATIS&#x2011;1000080 | STI certificates should contain a CertificatePolicies extension marked uncritical |
| [e_sti_subject_cn](../../ISSUES/e_sti_subject_cn/README.md) | error | ATIS&#x2011;1000080 | Common name shall contain the text string 'SHAKEN 8697' |
| [e_cp1_3_ambiguous_identifier](../../ISSUES/e_cp1_3_ambiguous_identifier/README.md) | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 31/10/2022 at 20:32:43