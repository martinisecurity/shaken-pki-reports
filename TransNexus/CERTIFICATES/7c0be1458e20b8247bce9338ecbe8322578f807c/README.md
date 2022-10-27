# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate 7c0be1458e20b8247bce9338ecbe8322578f807c
Tested At: 2022-10-27 18:56:20 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 182 day(s)\
Subject: CN=CCI SHAKEN 663J, OU=CCI, O=TransNexus, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA1, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.transnexus.com/663J/0b51e9f5-0e35-4d02-acef-5c91c3a6b903.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIICizCCAjGgAwIBAgIQYzbezCvctA5XQLmEW6N3WjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMTAeFw0yMjA0MjcxODIyNDdaFw0yMzA0MjcxODIyNDZaMEoxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpUcmFuc05leHVzMQwwCgYDVQQLEwNDQ0kxGDAWBgNVBAMTD0NDSSBTSEFLRU4gNjYzSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABMInL22gXrYYAUfgmDValv%2F13zNVsfRluSaD7yq6WHp0XkLJytYZNETa6xZ2YKq5Ucoq5aUkrtSsKIT%2BYDl37sejgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFADPCI5Ij1nKfDtw%2Fz4LjsVDkB58MB8GA1UdIwQYMBaAFJSGOT%2Fk9ZKgn%2F%2FcJ3UamkbweeFiMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFgYIKwYBBQUHARoECjAIoAYWBDY2M0owCgYIKoZIzj0EAwIDSAAwRQIgQoCURLu2mHjH%2FHASEbyslJIMXMKwMGCufT29%2FfJ49p0CIQD1HO1bL4BTNnxuMPiqKZfg%2Fga9ltO2Bj7QzD6tC%2BvRtg%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

Generated: 27/10/2022 at 18:57:27