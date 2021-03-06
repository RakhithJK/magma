#!/usr/bin/env python3
# Copyright (c) 2004-present Facebook All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

from typing import List, Optional, Tuple

from pysymphony import SymphonyClient

from .._utils import format_properties
from ..common.cache import LOCATION_TYPES
from ..common.data_class import Location, LocationType, PropertyValue
from ..common.data_enum import Entity
from ..exceptions import EntityNotFoundError
from ..graphql.input.add_location_type import AddLocationTypeInput
from ..graphql.mutation.add_location_type import AddLocationTypeMutation
from ..graphql.mutation.remove_location_type import RemoveLocationTypeMutation
from ..graphql.query.location_type_locations import LocationTypeLocationsQuery
from ..graphql.query.location_types import LocationTypesQuery
from .location import delete_location


def _populate_location_types(client: SymphonyClient) -> None:
    location_types = LocationTypesQuery.execute(client)
    if not location_types:
        return
    edges = location_types.edges
    for edge in edges:
        node = edge.node
        if node:
            LOCATION_TYPES[node.name] = LocationType(
                name=node.name, id=node.id, property_types=node.propertyTypes
            )


def add_location_type(
    client: SymphonyClient,
    name: str,
    properties: List[Tuple[str, str, Optional[PropertyValue], Optional[bool]]],
    map_zoom_level: int = 8,
) -> LocationType:
    """This function creates new location type.

        Args:
            name (str): location type name
            properties (List[Tuple[str, str, Optional[PropertyValue], Optional[bool]]]):
            - str - type name
            - str - enum["string", "int", "bool", "float", "date", "enum", "range",
            "email", "gps_location", "equipment", "location", "service", "datetime_local"]
            - PropertyValue - default property value
            - bool - fixed value flag

            map_zoom_level (int): map zoom level

        Returns:
            `pyinventory.common.data_class.LocationType` object

        Raises:
            FailedOperationException: internal inventory error

        Example:
            ```
            location_type = client.add_location_type(
                name="city",
                properties=[("Contact", "email", None, True)],
                map_zoom_level=5,
            )
            ```
    """
    new_property_types = format_properties(properties)
    result = AddLocationTypeMutation.execute(
        client,
        AddLocationTypeInput(
            name=name,
            mapZoomLevel=map_zoom_level,
            properties=new_property_types,
            surveyTemplateCategories=[],
        ),
    )

    location_type = LocationType(
        name=result.name, id=result.id, property_types=result.propertyTypes
    )
    LOCATION_TYPES[result.name] = location_type
    return location_type


def delete_locations_by_location_type(
    client: SymphonyClient, location_type: LocationType
) -> None:
    """Delete locatons by location type.

        Args:
            location_type ( `pyinventory.common.data_class.LocationType` ): location type object

        Raises:
            `pyinventory.exceptions.EntityNotFoundError`: if location_type does not exist

        Example:
            ```
            client.delete_locations_by_location_type(location_type=location_type)
            ```
    """
    location_type_with_locations = LocationTypeLocationsQuery.execute(
        client, id=location_type.id
    )
    if location_type_with_locations is None:
        raise EntityNotFoundError(
            entity=Entity.LocationType, entity_id=location_type.id
        )
    locations = location_type_with_locations.locations
    if locations is None:
        return
    for location in locations.edges:
        node = location.node
        if node:
            delete_location(
                client,
                Location(
                    id=node.id,
                    name=node.name,
                    latitude=node.latitude,
                    longitude=node.longitude,
                    externalId=node.externalId,
                    locationTypeName=node.locationType.name,
                ),
            )


def delete_location_type_with_locations(
    client: SymphonyClient, location_type: LocationType
) -> None:
    """Delete locaton type with existing locations.

        Args:
            location_type (`pyinventory.common.data_class.LocationType`): location type object

        Raises:
            `pyinventory.exceptions.EntityNotFoundError`: if location_type does not exist

        Example:
            ```
            client.delete_location_type_with_locations(location_type=location_type)
            ```
    """
    delete_locations_by_location_type(client, location_type)
    RemoveLocationTypeMutation.execute(client, id=location_type.id)
