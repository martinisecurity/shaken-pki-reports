# STIR/SHAKEN CA Ecosystem Compliance

## Certificate KEPS Technologies INC SHAKEN Cert 3535

Tested At: 02 Jun 23 01:00 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 428 day(s)\
Subject: CN=KEPS Technologies INC SHAKEN Cert 3535, O=KEPS Technologies INC, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/abaa4f0c25cc210facd07cd8c9e8d17192fdae3c

[View certificate details](https://understandingwebpki.com/?cert=MIICYzCCAgugAwIBAgIQEKPJPRla%2F%2BtkdBF6xVTYBTAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDgwMzEwNTcxNFoXDTI0MDgwMjEwNTcxNFowXjELMAkGA1UEBhMCVVMxHjAcBgNVBAoMFUtFUFMgVGVjaG5vbG9naWVzIElOQzEvMC0GA1UEAwwmS0VQUyBUZWNobm9sb2dpZXMgSU5DIFNIQUtFTiBDZXJ0IDM1MzUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ%2FdbmgGlfBonqvSEBUnV%2FcT%2FNwxLrKj7o1B48GjmOuolTvrZDz%2B6fEgXZK09r89WEiF6usoeSkLxxddnjOrjvAo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQzNTM1MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFD8d5Wzr%2Fl0gnDblMZhqW65e%2F97sMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0YAMEMCIFr6xXC0Ox6Eeex5lihOSsUuW1xhw15o9theUsao3zbUAh82pIxbSrQaXpBU68bzUK4Q4DbvZB42v%2BOBdAnIrI1o)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 02 Jun 23 01:12 UTC