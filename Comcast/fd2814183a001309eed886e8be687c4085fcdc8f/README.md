# STIR/SHAKEN CA Ecosystem Compliance
## Comcast

### Certificate fd2814183a001309eed886e8be687c4085fcdc8f
Tested At: 2022-10-26 20:32:16 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 8 day(s)\
Subject: CN=SHAKEN, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US

Link: https://sticr.stir.comcast.com/346670f0-c463-4c91-b0f3-8008a3f0102c.cer

View: [Click to view](https://understandingwebpki.com/?cert=MIICWDCCAf2gAwIBAgIIaCVar5XpRicwCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTIyMTAwNDEzMjk0N1oXDTIyMTEwMzEzMjk0N1owXjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MQ8wDQYDVQQDEwZTSEFLRU4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASLiT4QX7a%2FHO4oeF4U2P7kDMGm9YfUdHMmd29ZNU1xO8%2Frbf3CTIXgj8IQs4%2BTMKhtJ99QmP54gYeh9XH3CkrEo4GjMIGgMA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAWBggrBgEFBQcBGgQKMAigBhYENzYxMDAKBggqhkjOPQQDAgNJADBGAiEAtaJka%2BUzrfX8nJ9p9kJvC8H2jpLc%2FIjXwsDbuWgUaLsCIQCd5VifquZuJrEBI9VX1WcVXUvemKGTuB5DdjIs9IO22Q%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 7610' |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| w_ext_subject_key_identifier_missing_sub_cert | warn | RFC5280 |  |
| e_sti_subject_key_identifier | error | ATIS-1000080v4 | STI certificates shall contain a Subject Key Identifier extension |

Generated: 26/10/2022 at 20:32:17