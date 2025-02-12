#!/usr/bin/env python3

"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""

import ipaddress
import logging

from magma.common.misc_utils import get_ip_from_if_cidr
from magma.configuration.exceptions import LoadConfigError
from magma.configuration.mconfig_managers import load_service_mconfig_as_json
from magma.configuration.service_configs import load_service_config
from orc8r.protos.mconfig.mconfigs_pb2 import DnsD

from generate_service_config import generate_template_config

CONFIG_OVERRIDE_DIR = '/var/opt/magma/tmp'


def _get_addresses(cfg, mconfig):
    """
    Return list of addresses mapping domain to a type of record.

    EG: [{'ip': '192.88.99.142', 'domain': 'baiomc.cloudapp.net'}]}
    Currently uses record types A record, AAAA record and CNAME record.
    """
    # Start with list of addresses from YML
    addresses = cfg['addresses']
    for record in mconfig['records']:
        # Unpack each record type into list NOTE: list concat doesn't work here
        domain_records = [
            *record.a_record,
            *record.aaaa_record,
            #  *record.cname_record, TODO: Figure out how to repr CNAME
        ]
        for ip in domain_records:
            addresses.append({'domain': record.domain, 'ip': ip})

    return addresses


def get_context():
    """
    Provide context to pass to Jinja2 for templating.
    """
    context = {}
    cfg = load_service_config("dnsd")
    try:
        mconfig = load_service_mconfig_as_json("dnsd")
    except LoadConfigError as err:
        logging.warning("Error! Using default config because: %s", err)
        mconfig = DnsD()
    ip = get_ip_from_if_cidr(cfg['enodeb_interface'])
    dhcp_block_size = cfg['dhcp_block_size']
    available_hosts = list(ipaddress.IPv4Interface(ip).network.hosts())

    if dhcp_block_size < len(available_hosts):
        context['dhcp_range'] = {
            "lower": str(available_hosts[-dhcp_block_size]),
            "upper": str(available_hosts[-1]),
        }
    else:
        logging.fatal("Not enough available hosts to allocate a DHCP block of \
            %d addresses." % (dhcp_block_size))

    context['addresses'] = _get_addresses(cfg, mconfig)
    return context


def main():
    logging.basicConfig(
        level=logging.INFO,
        format='[%(asctime)s %(levelname)s %(name)s] %(message)s')

    generate_template_config('dnsd', 'dnsd',
                             CONFIG_OVERRIDE_DIR, get_context())


if __name__ == '__main__':
    main()
