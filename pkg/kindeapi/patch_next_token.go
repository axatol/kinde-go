package kindeapi

func (r GetAPIsResult) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetApplicationsResult) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetEnvironmentFeatureFlagsResult) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetOrganizationUsersResult) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetOrganizationsResult) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetOrganizationUserRolesResult) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetPermissionsResult) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetApplicationPropertyValuesResult) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetOrganizationPropertyValuesResult) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetUserPropertyValuesResult) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetRolesResult) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetSubscribersResult) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetRolePermissionsResult) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetUsersResult) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}
