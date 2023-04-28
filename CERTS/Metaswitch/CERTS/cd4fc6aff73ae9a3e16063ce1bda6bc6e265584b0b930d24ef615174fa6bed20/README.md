# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Nemont SHAKEN Cert 2247

Tested At: 28 Apr 23 02:16 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 388 day(s)\
Subject: CN=Nemont SHAKEN Cert 2247, O=Nemont, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/da7847867ee00785d849b849e374c81ecd3f2545

[View certificate details](https://understandingwebpki.com/?cert=MIICRjCCAe2gAwIBAgIQfnlUqz%2FjGhWuQpkRz9hSCDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDUyMDIyMTMzNFoXDTI0MDUxOTIyMTMzNFowQDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBk5lbW9udDEgMB4GA1UEAwwXTmVtb250IFNIQUtFTiBDZXJ0IDIyNDcwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAS%2BrXdXOqbl9jRCjhcjBTMbvnQSnqlU9G22qqqbWZjbrXdhMPPR7wiKayplO0VXvJ0JJvDBWCPozGRvvTAWBqt1o4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQyMjQ3MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFHEAsNJ6HPhOXQifjP39H1IxFXnqMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0cAMEQCIGOm5sWFQpVJRWQkydwwLEul7TNdnekR8c8CaWNVBatkAiBltUHKtrRAKeR%2FEIWrzkCnxkVyUd3V%2BHYZEwW3i3cufw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Apr 23 02:17 UTC