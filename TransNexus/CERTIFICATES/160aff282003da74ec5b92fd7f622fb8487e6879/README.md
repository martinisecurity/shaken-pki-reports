# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate 160aff282003da74ec5b92fd7f622fb8487e6879
Tested At: 2022-10-28 10:32:05 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 28 day(s)\
Subject: CN=SHAKEN 722J, OU=SHAKEN, O=EvolveIP LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.transnexus.com/722J/f75ff50a-1fd7-4087-a0a1-de57acb1afc5.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC7jCCApSgAwIBAgIQc%2BKyCHzJSNGtnXxVDfceGjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjUxNTMzMDhaFw0yMjExMjQxNTMzMDdaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwxFdm9sdmVJUCBMTEMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDcyMkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ8f0w80%2F9ZFljqzq0go21KcrjLZ7qmhEyg45SEWPAG%2F8Fo2Jn%2FYAenoRoCzw7IuqU8P6QUG5Fq4YXyjxMk5otUo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFAI7mT3J%2FhAQ7DXaxwt9jbd%2BiZnqMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDcyMkowCgYIKoZIzj0EAwIDSAAwRQIhANIfcvfuWaXSLlHKydwmC%2Be%2Bq5R1X2vvW5S2B4GH8wCGAiBMsL5Lxw4VDpM7oxcdjw8%2BO%2Fu8lHlpgBF8BEIl3%2Bd0DA%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 28/10/2022 at 10:33:25