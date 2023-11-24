# STIR/SHAKEN CA Ecosystem Compliance

## NetNumber

Name: e_atis_ext_crl_distribution_struct\
Source: ATIS1000080\
Citation: ATIS-1000080.v004 / 6.4.1 STI Certificate Requirements\
Effective Date: 16 Jan 22 00:00 UTC\
Description: STI intermediate and end-entity certificates shall contain a CRL Distribution Points extension containing a single DistributionPoint entry. The DistributionPoint entry shall contain a distributionPoint field identifying the HTTP URL reference to the file containing the SHAKEN CRL hosted by the STI-PA, and a cRLIssuer field that contains the DN of the issuer of the CRL.

### Leaf Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| error | Plivo Inc | [view](../../CERTS/fed50200daa631dd0cd7b74969c780f8d456dfd31db156c6cbea276f5a9a4cbf/README.md) | CRL Distribution Point shall contain a CRLIssuer field |
| error | DISH Wireless L.L.C.SHAKEN.490J | [view](../../CERTS/2943713c56f0705ed027ecffced5eb89cb1c36bb5386bdc36a6b8e5618ca2c9c/README.md) | CRL Distribution Point shall contain a CRLIssuer field |
| error | Google SHAKEN cert 969H | [view](../../CERTS/16f320d1971e6da38e8a26433b6d8006ff2144912ff1f128c51da37cfc2bd6c3/README.md) | CRL Distribution Point shall contain a CRLIssuer field |

### CA Certificates

No error, warning, or notice level issues were found


Generated: 24 Nov 23 11:17 UTC