# STIR/SHAKEN CA Ecosystem Compliance

## Telonium

Name: e_atis_ext_crl_distribution_struct_ca\
Source: ATIS1000080\
Citation: ATIS-1000080.v004 / 6.4.1 STI Certificate Requirements\
Effective Date: 18 Oct 21 00:00 UTC\
Description: STI intermediate and end-entity certificates shall contain a CRL Distribution Points extension containing a single DistributionPoint entry. The DistributionPoint entry shall contain a distributionPoint field identifying the HTTP URL reference to the file containing the SHAKEN CRL hosted by the STI-PA, and a cRLIssuer field that contains the DN of the issuer of the CRL.

### Leaf Certificates

No error, warning, or notice level issues were found

### CA Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| error | Telonium SHAKEN Intermediate G1 | [view](../../CERTS/8c0ef8682826bec79a8c64881899f6a5a4a1d52dfebe28ae419c23f85df96ea0/README.md) | CRL Distribution Point shall contain a CRLIssuer field |


Generated: 24 Nov 23 11:17 UTC