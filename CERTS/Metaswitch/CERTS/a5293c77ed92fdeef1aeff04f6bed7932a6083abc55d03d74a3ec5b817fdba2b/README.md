# STIR/SHAKEN CA Ecosystem Compliance

## Certificate ENA SHAKEN cert 521F

Tested At: 16 Mar 23 19:13 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 456 day(s)\
Subject: CN=ENA SHAKEN cert 521F, O=ENA, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/702fb83e8268b5175b591d8448a512646339bc40

[View certificate details](https://understandingwebpki.com/?cert=MIICQTCCAeegAwIBAgIQH2O8LTasSJz0F8wWympA3TAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDYxNTEzNTY1MVoXDTI0MDYxNDEzNTY1MVowOjELMAkGA1UEBhMCVVMxDDAKBgNVBAoMA0VOQTEdMBsGA1UEAwwURU5BIFNIQUtFTiBjZXJ0IDUyMUYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATww3ygb81IGEdAyq4Dc0OJEUAOF3kLhGB9HJF0xOihIvMmHQqpO8v95IEi0tt8usB7%2Bx2w0FadPsBYjzPZ%2F6ylo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ1MjFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFHkOd0WGEQe5gWecn95SH%2FvhOqbnMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQDmdX3vvawVcYCD0dgkIjqmKcaa2wD%2FUfNIEe31fAWtCwIgaYXGNP25PAhY9k%2F3tlyHAtTlETsEqViAVFECul%2FuuLE%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 16 Mar 23 19:18 UTC