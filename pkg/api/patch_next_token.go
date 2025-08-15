package api

func (r GetAPIsResp) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetApplicationsResp) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetEnvironmentFeatureFlagsResp) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetOrganizationUsersResp) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetOrganizationsResp) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetOrganizationUserRolesResp) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetPermissionsResp) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetApplicationPropertyValuesResp) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetOrganizationPropertyValuesResp) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetUserPropertyValuesResp) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetRolesResp) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetSubscribersResp) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetRolePermissionsResp) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}

func (r GetUsersResp) NextToken() *string {
	if r.JSON200 != nil && r.JSON200.NextToken != nil {
		return r.JSON200.NextToken
	}
	return nil
}
