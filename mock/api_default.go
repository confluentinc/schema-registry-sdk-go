// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: api_default.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_schema_registry_sdk_go "github.com/confluentinc/schema-registry-sdk-go"
)

// DefaultApi is a mock of DefaultApi interface
type DefaultApi struct {
	lockDeleteSchemaVersion sync.Mutex
	DeleteSchemaVersionFunc func(ctx context.Context, subject, version string) (int32, *net_http.Response, error)

	lockDeleteSubject sync.Mutex
	DeleteSubjectFunc func(ctx context.Context, subject string) ([]int32, *net_http.Response, error)

	lockGet sync.Mutex
	GetFunc func(ctx context.Context) (map[string]map[string]interface{}, *net_http.Response, error)

	lockGetMode sync.Mutex
	GetModeFunc func(ctx context.Context, subject string) (github_com_confluentinc_schema_registry_sdk_go.ModeGetResponse, *net_http.Response, error)

	lockGetSchema sync.Mutex
	GetSchemaFunc func(ctx context.Context, id int32) (github_com_confluentinc_schema_registry_sdk_go.SchemaString, *net_http.Response, error)

	lockGetSchemaByVersion sync.Mutex
	GetSchemaByVersionFunc func(ctx context.Context, subject, version string) (github_com_confluentinc_schema_registry_sdk_go.Schema, *net_http.Response, error)

	lockGetSchemaOnly sync.Mutex
	GetSchemaOnlyFunc func(ctx context.Context, subject, version string) (string, *net_http.Response, error)

	lockGetSubjectLevelConfig sync.Mutex
	GetSubjectLevelConfigFunc func(ctx context.Context, subject string) (github_com_confluentinc_schema_registry_sdk_go.Config, *net_http.Response, error)

	lockGetTopLevelConfig sync.Mutex
	GetTopLevelConfigFunc func(ctx context.Context) (github_com_confluentinc_schema_registry_sdk_go.Config, *net_http.Response, error)

	lockGetTopLevelMode sync.Mutex
	GetTopLevelModeFunc func(ctx context.Context) (github_com_confluentinc_schema_registry_sdk_go.ModeGetResponse, *net_http.Response, error)

	lockList sync.Mutex
	ListFunc func(ctx context.Context) ([]string, *net_http.Response, error)

	lockListVersions sync.Mutex
	ListVersionsFunc func(ctx context.Context, subject string) ([]int32, *net_http.Response, error)

	lockLookUpSchemaUnderSubject sync.Mutex
	LookUpSchemaUnderSubjectFunc func(ctx context.Context, subject string, body github_com_confluentinc_schema_registry_sdk_go.RegisterSchemaRequest, localVarOptionals *github_com_confluentinc_schema_registry_sdk_go.LookUpSchemaUnderSubjectOpts) (*net_http.Response, error)

	lockPost sync.Mutex
	PostFunc func(ctx context.Context) (map[string]string, *net_http.Response, error)

	lockRegister sync.Mutex
	RegisterFunc func(ctx context.Context, subject string, body github_com_confluentinc_schema_registry_sdk_go.RegisterSchemaRequest) (github_com_confluentinc_schema_registry_sdk_go.RegisterSchemaResponse, *net_http.Response, error)

	lockTestCompatabilityBySubjectName sync.Mutex
	TestCompatabilityBySubjectNameFunc func(ctx context.Context, subject, version string, body github_com_confluentinc_schema_registry_sdk_go.RegisterSchemaRequest, localVarOptionals *github_com_confluentinc_schema_registry_sdk_go.TestCompatabilityBySubjectNameOpts) (github_com_confluentinc_schema_registry_sdk_go.CompatibilityCheckResponse, *net_http.Response, error)

	lockUpdateMode sync.Mutex
	UpdateModeFunc func(ctx context.Context, subject string) (github_com_confluentinc_schema_registry_sdk_go.ModeUpdateRequest, *net_http.Response, error)

	lockUpdateSubjectLevelConfig sync.Mutex
	UpdateSubjectLevelConfigFunc func(ctx context.Context, subject string, body github_com_confluentinc_schema_registry_sdk_go.ConfigUpdateRequest) (github_com_confluentinc_schema_registry_sdk_go.ConfigUpdateRequest, *net_http.Response, error)

	lockUpdateTopLevelConfig sync.Mutex
	UpdateTopLevelConfigFunc func(ctx context.Context, body github_com_confluentinc_schema_registry_sdk_go.ConfigUpdateRequest) (github_com_confluentinc_schema_registry_sdk_go.ConfigUpdateRequest, *net_http.Response, error)

	lockUpdateTopLevelMode sync.Mutex
	UpdateTopLevelModeFunc func(ctx context.Context) (github_com_confluentinc_schema_registry_sdk_go.ModeUpdateRequest, *net_http.Response, error)

	calls struct {
		DeleteSchemaVersion []struct {
			Ctx     context.Context
			Subject string
			Version string
		}
		DeleteSubject []struct {
			Ctx     context.Context
			Subject string
		}
		Get []struct {
			Ctx context.Context
		}
		GetMode []struct {
			Ctx     context.Context
			Subject string
		}
		GetSchema []struct {
			Ctx context.Context
			Id  int32
		}
		GetSchemaByVersion []struct {
			Ctx     context.Context
			Subject string
			Version string
		}
		GetSchemaOnly []struct {
			Ctx     context.Context
			Subject string
			Version string
		}
		GetSubjectLevelConfig []struct {
			Ctx     context.Context
			Subject string
		}
		GetTopLevelConfig []struct {
			Ctx context.Context
		}
		GetTopLevelMode []struct {
			Ctx context.Context
		}
		List []struct {
			Ctx context.Context
		}
		ListVersions []struct {
			Ctx     context.Context
			Subject string
		}
		LookUpSchemaUnderSubject []struct {
			Ctx               context.Context
			Subject           string
			Body              github_com_confluentinc_schema_registry_sdk_go.RegisterSchemaRequest
			LocalVarOptionals *github_com_confluentinc_schema_registry_sdk_go.LookUpSchemaUnderSubjectOpts
		}
		Post []struct {
			Ctx context.Context
		}
		Register []struct {
			Ctx     context.Context
			Subject string
			Body    github_com_confluentinc_schema_registry_sdk_go.RegisterSchemaRequest
		}
		TestCompatabilityBySubjectName []struct {
			Ctx               context.Context
			Subject           string
			Version           string
			Body              github_com_confluentinc_schema_registry_sdk_go.RegisterSchemaRequest
			LocalVarOptionals *github_com_confluentinc_schema_registry_sdk_go.TestCompatabilityBySubjectNameOpts
		}
		UpdateMode []struct {
			Ctx     context.Context
			Subject string
		}
		UpdateSubjectLevelConfig []struct {
			Ctx     context.Context
			Subject string
			Body    github_com_confluentinc_schema_registry_sdk_go.ConfigUpdateRequest
		}
		UpdateTopLevelConfig []struct {
			Ctx  context.Context
			Body github_com_confluentinc_schema_registry_sdk_go.ConfigUpdateRequest
		}
		UpdateTopLevelMode []struct {
			Ctx context.Context
		}
	}
}

// DeleteSchemaVersion mocks base method by wrapping the associated func.
func (m *DefaultApi) DeleteSchemaVersion(ctx context.Context, subject, version string) (int32, *net_http.Response, error) {
	m.lockDeleteSchemaVersion.Lock()
	defer m.lockDeleteSchemaVersion.Unlock()

	if m.DeleteSchemaVersionFunc == nil {
		panic("mocker: DefaultApi.DeleteSchemaVersionFunc is nil but DefaultApi.DeleteSchemaVersion was called.")
	}

	call := struct {
		Ctx     context.Context
		Subject string
		Version string
	}{
		Ctx:     ctx,
		Subject: subject,
		Version: version,
	}

	m.calls.DeleteSchemaVersion = append(m.calls.DeleteSchemaVersion, call)

	return m.DeleteSchemaVersionFunc(ctx, subject, version)
}

// DeleteSchemaVersionCalled returns true if DeleteSchemaVersion was called at least once.
func (m *DefaultApi) DeleteSchemaVersionCalled() bool {
	m.lockDeleteSchemaVersion.Lock()
	defer m.lockDeleteSchemaVersion.Unlock()

	return len(m.calls.DeleteSchemaVersion) > 0
}

// DeleteSchemaVersionCalls returns the calls made to DeleteSchemaVersion.
func (m *DefaultApi) DeleteSchemaVersionCalls() []struct {
	Ctx     context.Context
	Subject string
	Version string
} {
	m.lockDeleteSchemaVersion.Lock()
	defer m.lockDeleteSchemaVersion.Unlock()

	return m.calls.DeleteSchemaVersion
}

// DeleteSubject mocks base method by wrapping the associated func.
func (m *DefaultApi) DeleteSubject(ctx context.Context, subject string) ([]int32, *net_http.Response, error) {
	m.lockDeleteSubject.Lock()
	defer m.lockDeleteSubject.Unlock()

	if m.DeleteSubjectFunc == nil {
		panic("mocker: DefaultApi.DeleteSubjectFunc is nil but DefaultApi.DeleteSubject was called.")
	}

	call := struct {
		Ctx     context.Context
		Subject string
	}{
		Ctx:     ctx,
		Subject: subject,
	}

	m.calls.DeleteSubject = append(m.calls.DeleteSubject, call)

	return m.DeleteSubjectFunc(ctx, subject)
}

// DeleteSubjectCalled returns true if DeleteSubject was called at least once.
func (m *DefaultApi) DeleteSubjectCalled() bool {
	m.lockDeleteSubject.Lock()
	defer m.lockDeleteSubject.Unlock()

	return len(m.calls.DeleteSubject) > 0
}

// DeleteSubjectCalls returns the calls made to DeleteSubject.
func (m *DefaultApi) DeleteSubjectCalls() []struct {
	Ctx     context.Context
	Subject string
} {
	m.lockDeleteSubject.Lock()
	defer m.lockDeleteSubject.Unlock()

	return m.calls.DeleteSubject
}

// Get mocks base method by wrapping the associated func.
func (m *DefaultApi) Get(ctx context.Context) (map[string]map[string]interface{}, *net_http.Response, error) {
	m.lockGet.Lock()
	defer m.lockGet.Unlock()

	if m.GetFunc == nil {
		panic("mocker: DefaultApi.GetFunc is nil but DefaultApi.Get was called.")
	}

	call := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}

	m.calls.Get = append(m.calls.Get, call)

	return m.GetFunc(ctx)
}

// GetCalled returns true if Get was called at least once.
func (m *DefaultApi) GetCalled() bool {
	m.lockGet.Lock()
	defer m.lockGet.Unlock()

	return len(m.calls.Get) > 0
}

// GetCalls returns the calls made to Get.
func (m *DefaultApi) GetCalls() []struct {
	Ctx context.Context
} {
	m.lockGet.Lock()
	defer m.lockGet.Unlock()

	return m.calls.Get
}

// GetMode mocks base method by wrapping the associated func.
func (m *DefaultApi) GetMode(ctx context.Context, subject string) (github_com_confluentinc_schema_registry_sdk_go.ModeGetResponse, *net_http.Response, error) {
	m.lockGetMode.Lock()
	defer m.lockGetMode.Unlock()

	if m.GetModeFunc == nil {
		panic("mocker: DefaultApi.GetModeFunc is nil but DefaultApi.GetMode was called.")
	}

	call := struct {
		Ctx     context.Context
		Subject string
	}{
		Ctx:     ctx,
		Subject: subject,
	}

	m.calls.GetMode = append(m.calls.GetMode, call)

	return m.GetModeFunc(ctx, subject)
}

// GetModeCalled returns true if GetMode was called at least once.
func (m *DefaultApi) GetModeCalled() bool {
	m.lockGetMode.Lock()
	defer m.lockGetMode.Unlock()

	return len(m.calls.GetMode) > 0
}

// GetModeCalls returns the calls made to GetMode.
func (m *DefaultApi) GetModeCalls() []struct {
	Ctx     context.Context
	Subject string
} {
	m.lockGetMode.Lock()
	defer m.lockGetMode.Unlock()

	return m.calls.GetMode
}

// GetSchema mocks base method by wrapping the associated func.
func (m *DefaultApi) GetSchema(ctx context.Context, id int32) (github_com_confluentinc_schema_registry_sdk_go.SchemaString, *net_http.Response, error) {
	m.lockGetSchema.Lock()
	defer m.lockGetSchema.Unlock()

	if m.GetSchemaFunc == nil {
		panic("mocker: DefaultApi.GetSchemaFunc is nil but DefaultApi.GetSchema was called.")
	}

	call := struct {
		Ctx context.Context
		Id  int32
	}{
		Ctx: ctx,
		Id:  id,
	}

	m.calls.GetSchema = append(m.calls.GetSchema, call)

	return m.GetSchemaFunc(ctx, id)
}

// GetSchemaCalled returns true if GetSchema was called at least once.
func (m *DefaultApi) GetSchemaCalled() bool {
	m.lockGetSchema.Lock()
	defer m.lockGetSchema.Unlock()

	return len(m.calls.GetSchema) > 0
}

// GetSchemaCalls returns the calls made to GetSchema.
func (m *DefaultApi) GetSchemaCalls() []struct {
	Ctx context.Context
	Id  int32
} {
	m.lockGetSchema.Lock()
	defer m.lockGetSchema.Unlock()

	return m.calls.GetSchema
}

// GetSchemaByVersion mocks base method by wrapping the associated func.
func (m *DefaultApi) GetSchemaByVersion(ctx context.Context, subject, version string) (github_com_confluentinc_schema_registry_sdk_go.Schema, *net_http.Response, error) {
	m.lockGetSchemaByVersion.Lock()
	defer m.lockGetSchemaByVersion.Unlock()

	if m.GetSchemaByVersionFunc == nil {
		panic("mocker: DefaultApi.GetSchemaByVersionFunc is nil but DefaultApi.GetSchemaByVersion was called.")
	}

	call := struct {
		Ctx     context.Context
		Subject string
		Version string
	}{
		Ctx:     ctx,
		Subject: subject,
		Version: version,
	}

	m.calls.GetSchemaByVersion = append(m.calls.GetSchemaByVersion, call)

	return m.GetSchemaByVersionFunc(ctx, subject, version)
}

// GetSchemaByVersionCalled returns true if GetSchemaByVersion was called at least once.
func (m *DefaultApi) GetSchemaByVersionCalled() bool {
	m.lockGetSchemaByVersion.Lock()
	defer m.lockGetSchemaByVersion.Unlock()

	return len(m.calls.GetSchemaByVersion) > 0
}

// GetSchemaByVersionCalls returns the calls made to GetSchemaByVersion.
func (m *DefaultApi) GetSchemaByVersionCalls() []struct {
	Ctx     context.Context
	Subject string
	Version string
} {
	m.lockGetSchemaByVersion.Lock()
	defer m.lockGetSchemaByVersion.Unlock()

	return m.calls.GetSchemaByVersion
}

// GetSchemaOnly mocks base method by wrapping the associated func.
func (m *DefaultApi) GetSchemaOnly(ctx context.Context, subject, version string) (string, *net_http.Response, error) {
	m.lockGetSchemaOnly.Lock()
	defer m.lockGetSchemaOnly.Unlock()

	if m.GetSchemaOnlyFunc == nil {
		panic("mocker: DefaultApi.GetSchemaOnlyFunc is nil but DefaultApi.GetSchemaOnly was called.")
	}

	call := struct {
		Ctx     context.Context
		Subject string
		Version string
	}{
		Ctx:     ctx,
		Subject: subject,
		Version: version,
	}

	m.calls.GetSchemaOnly = append(m.calls.GetSchemaOnly, call)

	return m.GetSchemaOnlyFunc(ctx, subject, version)
}

// GetSchemaOnlyCalled returns true if GetSchemaOnly was called at least once.
func (m *DefaultApi) GetSchemaOnlyCalled() bool {
	m.lockGetSchemaOnly.Lock()
	defer m.lockGetSchemaOnly.Unlock()

	return len(m.calls.GetSchemaOnly) > 0
}

// GetSchemaOnlyCalls returns the calls made to GetSchemaOnly.
func (m *DefaultApi) GetSchemaOnlyCalls() []struct {
	Ctx     context.Context
	Subject string
	Version string
} {
	m.lockGetSchemaOnly.Lock()
	defer m.lockGetSchemaOnly.Unlock()

	return m.calls.GetSchemaOnly
}

// GetSubjectLevelConfig mocks base method by wrapping the associated func.
func (m *DefaultApi) GetSubjectLevelConfig(ctx context.Context, subject string) (github_com_confluentinc_schema_registry_sdk_go.Config, *net_http.Response, error) {
	m.lockGetSubjectLevelConfig.Lock()
	defer m.lockGetSubjectLevelConfig.Unlock()

	if m.GetSubjectLevelConfigFunc == nil {
		panic("mocker: DefaultApi.GetSubjectLevelConfigFunc is nil but DefaultApi.GetSubjectLevelConfig was called.")
	}

	call := struct {
		Ctx     context.Context
		Subject string
	}{
		Ctx:     ctx,
		Subject: subject,
	}

	m.calls.GetSubjectLevelConfig = append(m.calls.GetSubjectLevelConfig, call)

	return m.GetSubjectLevelConfigFunc(ctx, subject)
}

// GetSubjectLevelConfigCalled returns true if GetSubjectLevelConfig was called at least once.
func (m *DefaultApi) GetSubjectLevelConfigCalled() bool {
	m.lockGetSubjectLevelConfig.Lock()
	defer m.lockGetSubjectLevelConfig.Unlock()

	return len(m.calls.GetSubjectLevelConfig) > 0
}

// GetSubjectLevelConfigCalls returns the calls made to GetSubjectLevelConfig.
func (m *DefaultApi) GetSubjectLevelConfigCalls() []struct {
	Ctx     context.Context
	Subject string
} {
	m.lockGetSubjectLevelConfig.Lock()
	defer m.lockGetSubjectLevelConfig.Unlock()

	return m.calls.GetSubjectLevelConfig
}

// GetTopLevelConfig mocks base method by wrapping the associated func.
func (m *DefaultApi) GetTopLevelConfig(ctx context.Context) (github_com_confluentinc_schema_registry_sdk_go.Config, *net_http.Response, error) {
	m.lockGetTopLevelConfig.Lock()
	defer m.lockGetTopLevelConfig.Unlock()

	if m.GetTopLevelConfigFunc == nil {
		panic("mocker: DefaultApi.GetTopLevelConfigFunc is nil but DefaultApi.GetTopLevelConfig was called.")
	}

	call := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}

	m.calls.GetTopLevelConfig = append(m.calls.GetTopLevelConfig, call)

	return m.GetTopLevelConfigFunc(ctx)
}

// GetTopLevelConfigCalled returns true if GetTopLevelConfig was called at least once.
func (m *DefaultApi) GetTopLevelConfigCalled() bool {
	m.lockGetTopLevelConfig.Lock()
	defer m.lockGetTopLevelConfig.Unlock()

	return len(m.calls.GetTopLevelConfig) > 0
}

// GetTopLevelConfigCalls returns the calls made to GetTopLevelConfig.
func (m *DefaultApi) GetTopLevelConfigCalls() []struct {
	Ctx context.Context
} {
	m.lockGetTopLevelConfig.Lock()
	defer m.lockGetTopLevelConfig.Unlock()

	return m.calls.GetTopLevelConfig
}

// GetTopLevelMode mocks base method by wrapping the associated func.
func (m *DefaultApi) GetTopLevelMode(ctx context.Context) (github_com_confluentinc_schema_registry_sdk_go.ModeGetResponse, *net_http.Response, error) {
	m.lockGetTopLevelMode.Lock()
	defer m.lockGetTopLevelMode.Unlock()

	if m.GetTopLevelModeFunc == nil {
		panic("mocker: DefaultApi.GetTopLevelModeFunc is nil but DefaultApi.GetTopLevelMode was called.")
	}

	call := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}

	m.calls.GetTopLevelMode = append(m.calls.GetTopLevelMode, call)

	return m.GetTopLevelModeFunc(ctx)
}

// GetTopLevelModeCalled returns true if GetTopLevelMode was called at least once.
func (m *DefaultApi) GetTopLevelModeCalled() bool {
	m.lockGetTopLevelMode.Lock()
	defer m.lockGetTopLevelMode.Unlock()

	return len(m.calls.GetTopLevelMode) > 0
}

// GetTopLevelModeCalls returns the calls made to GetTopLevelMode.
func (m *DefaultApi) GetTopLevelModeCalls() []struct {
	Ctx context.Context
} {
	m.lockGetTopLevelMode.Lock()
	defer m.lockGetTopLevelMode.Unlock()

	return m.calls.GetTopLevelMode
}

// List mocks base method by wrapping the associated func.
func (m *DefaultApi) List(ctx context.Context) ([]string, *net_http.Response, error) {
	m.lockList.Lock()
	defer m.lockList.Unlock()

	if m.ListFunc == nil {
		panic("mocker: DefaultApi.ListFunc is nil but DefaultApi.List was called.")
	}

	call := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}

	m.calls.List = append(m.calls.List, call)

	return m.ListFunc(ctx)
}

// ListCalled returns true if List was called at least once.
func (m *DefaultApi) ListCalled() bool {
	m.lockList.Lock()
	defer m.lockList.Unlock()

	return len(m.calls.List) > 0
}

// ListCalls returns the calls made to List.
func (m *DefaultApi) ListCalls() []struct {
	Ctx context.Context
} {
	m.lockList.Lock()
	defer m.lockList.Unlock()

	return m.calls.List
}

// ListVersions mocks base method by wrapping the associated func.
func (m *DefaultApi) ListVersions(ctx context.Context, subject string) ([]int32, *net_http.Response, error) {
	m.lockListVersions.Lock()
	defer m.lockListVersions.Unlock()

	if m.ListVersionsFunc == nil {
		panic("mocker: DefaultApi.ListVersionsFunc is nil but DefaultApi.ListVersions was called.")
	}

	call := struct {
		Ctx     context.Context
		Subject string
	}{
		Ctx:     ctx,
		Subject: subject,
	}

	m.calls.ListVersions = append(m.calls.ListVersions, call)

	return m.ListVersionsFunc(ctx, subject)
}

// ListVersionsCalled returns true if ListVersions was called at least once.
func (m *DefaultApi) ListVersionsCalled() bool {
	m.lockListVersions.Lock()
	defer m.lockListVersions.Unlock()

	return len(m.calls.ListVersions) > 0
}

// ListVersionsCalls returns the calls made to ListVersions.
func (m *DefaultApi) ListVersionsCalls() []struct {
	Ctx     context.Context
	Subject string
} {
	m.lockListVersions.Lock()
	defer m.lockListVersions.Unlock()

	return m.calls.ListVersions
}

// LookUpSchemaUnderSubject mocks base method by wrapping the associated func.
func (m *DefaultApi) LookUpSchemaUnderSubject(ctx context.Context, subject string, body github_com_confluentinc_schema_registry_sdk_go.RegisterSchemaRequest, localVarOptionals *github_com_confluentinc_schema_registry_sdk_go.LookUpSchemaUnderSubjectOpts) (*net_http.Response, error) {
	m.lockLookUpSchemaUnderSubject.Lock()
	defer m.lockLookUpSchemaUnderSubject.Unlock()

	if m.LookUpSchemaUnderSubjectFunc == nil {
		panic("mocker: DefaultApi.LookUpSchemaUnderSubjectFunc is nil but DefaultApi.LookUpSchemaUnderSubject was called.")
	}

	call := struct {
		Ctx               context.Context
		Subject           string
		Body              github_com_confluentinc_schema_registry_sdk_go.RegisterSchemaRequest
		LocalVarOptionals *github_com_confluentinc_schema_registry_sdk_go.LookUpSchemaUnderSubjectOpts
	}{
		Ctx:               ctx,
		Subject:           subject,
		Body:              body,
		LocalVarOptionals: localVarOptionals,
	}

	m.calls.LookUpSchemaUnderSubject = append(m.calls.LookUpSchemaUnderSubject, call)

	return m.LookUpSchemaUnderSubjectFunc(ctx, subject, body, localVarOptionals)
}

// LookUpSchemaUnderSubjectCalled returns true if LookUpSchemaUnderSubject was called at least once.
func (m *DefaultApi) LookUpSchemaUnderSubjectCalled() bool {
	m.lockLookUpSchemaUnderSubject.Lock()
	defer m.lockLookUpSchemaUnderSubject.Unlock()

	return len(m.calls.LookUpSchemaUnderSubject) > 0
}

// LookUpSchemaUnderSubjectCalls returns the calls made to LookUpSchemaUnderSubject.
func (m *DefaultApi) LookUpSchemaUnderSubjectCalls() []struct {
	Ctx               context.Context
	Subject           string
	Body              github_com_confluentinc_schema_registry_sdk_go.RegisterSchemaRequest
	LocalVarOptionals *github_com_confluentinc_schema_registry_sdk_go.LookUpSchemaUnderSubjectOpts
} {
	m.lockLookUpSchemaUnderSubject.Lock()
	defer m.lockLookUpSchemaUnderSubject.Unlock()

	return m.calls.LookUpSchemaUnderSubject
}

// Post mocks base method by wrapping the associated func.
func (m *DefaultApi) Post(ctx context.Context) (map[string]string, *net_http.Response, error) {
	m.lockPost.Lock()
	defer m.lockPost.Unlock()

	if m.PostFunc == nil {
		panic("mocker: DefaultApi.PostFunc is nil but DefaultApi.Post was called.")
	}

	call := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}

	m.calls.Post = append(m.calls.Post, call)

	return m.PostFunc(ctx)
}

// PostCalled returns true if Post was called at least once.
func (m *DefaultApi) PostCalled() bool {
	m.lockPost.Lock()
	defer m.lockPost.Unlock()

	return len(m.calls.Post) > 0
}

// PostCalls returns the calls made to Post.
func (m *DefaultApi) PostCalls() []struct {
	Ctx context.Context
} {
	m.lockPost.Lock()
	defer m.lockPost.Unlock()

	return m.calls.Post
}

// Register mocks base method by wrapping the associated func.
func (m *DefaultApi) Register(ctx context.Context, subject string, body github_com_confluentinc_schema_registry_sdk_go.RegisterSchemaRequest) (github_com_confluentinc_schema_registry_sdk_go.RegisterSchemaResponse, *net_http.Response, error) {
	m.lockRegister.Lock()
	defer m.lockRegister.Unlock()

	if m.RegisterFunc == nil {
		panic("mocker: DefaultApi.RegisterFunc is nil but DefaultApi.Register was called.")
	}

	call := struct {
		Ctx     context.Context
		Subject string
		Body    github_com_confluentinc_schema_registry_sdk_go.RegisterSchemaRequest
	}{
		Ctx:     ctx,
		Subject: subject,
		Body:    body,
	}

	m.calls.Register = append(m.calls.Register, call)

	return m.RegisterFunc(ctx, subject, body)
}

// RegisterCalled returns true if Register was called at least once.
func (m *DefaultApi) RegisterCalled() bool {
	m.lockRegister.Lock()
	defer m.lockRegister.Unlock()

	return len(m.calls.Register) > 0
}

// RegisterCalls returns the calls made to Register.
func (m *DefaultApi) RegisterCalls() []struct {
	Ctx     context.Context
	Subject string
	Body    github_com_confluentinc_schema_registry_sdk_go.RegisterSchemaRequest
} {
	m.lockRegister.Lock()
	defer m.lockRegister.Unlock()

	return m.calls.Register
}

// TestCompatabilityBySubjectName mocks base method by wrapping the associated func.
func (m *DefaultApi) TestCompatabilityBySubjectName(ctx context.Context, subject, version string, body github_com_confluentinc_schema_registry_sdk_go.RegisterSchemaRequest, localVarOptionals *github_com_confluentinc_schema_registry_sdk_go.TestCompatabilityBySubjectNameOpts) (github_com_confluentinc_schema_registry_sdk_go.CompatibilityCheckResponse, *net_http.Response, error) {
	m.lockTestCompatabilityBySubjectName.Lock()
	defer m.lockTestCompatabilityBySubjectName.Unlock()

	if m.TestCompatabilityBySubjectNameFunc == nil {
		panic("mocker: DefaultApi.TestCompatabilityBySubjectNameFunc is nil but DefaultApi.TestCompatabilityBySubjectName was called.")
	}

	call := struct {
		Ctx               context.Context
		Subject           string
		Version           string
		Body              github_com_confluentinc_schema_registry_sdk_go.RegisterSchemaRequest
		LocalVarOptionals *github_com_confluentinc_schema_registry_sdk_go.TestCompatabilityBySubjectNameOpts
	}{
		Ctx:               ctx,
		Subject:           subject,
		Version:           version,
		Body:              body,
		LocalVarOptionals: localVarOptionals,
	}

	m.calls.TestCompatabilityBySubjectName = append(m.calls.TestCompatabilityBySubjectName, call)

	return m.TestCompatabilityBySubjectNameFunc(ctx, subject, version, body, localVarOptionals)
}

// TestCompatabilityBySubjectNameCalled returns true if TestCompatabilityBySubjectName was called at least once.
func (m *DefaultApi) TestCompatabilityBySubjectNameCalled() bool {
	m.lockTestCompatabilityBySubjectName.Lock()
	defer m.lockTestCompatabilityBySubjectName.Unlock()

	return len(m.calls.TestCompatabilityBySubjectName) > 0
}

// TestCompatabilityBySubjectNameCalls returns the calls made to TestCompatabilityBySubjectName.
func (m *DefaultApi) TestCompatabilityBySubjectNameCalls() []struct {
	Ctx               context.Context
	Subject           string
	Version           string
	Body              github_com_confluentinc_schema_registry_sdk_go.RegisterSchemaRequest
	LocalVarOptionals *github_com_confluentinc_schema_registry_sdk_go.TestCompatabilityBySubjectNameOpts
} {
	m.lockTestCompatabilityBySubjectName.Lock()
	defer m.lockTestCompatabilityBySubjectName.Unlock()

	return m.calls.TestCompatabilityBySubjectName
}

// UpdateMode mocks base method by wrapping the associated func.
func (m *DefaultApi) UpdateMode(ctx context.Context, subject string) (github_com_confluentinc_schema_registry_sdk_go.ModeUpdateRequest, *net_http.Response, error) {
	m.lockUpdateMode.Lock()
	defer m.lockUpdateMode.Unlock()

	if m.UpdateModeFunc == nil {
		panic("mocker: DefaultApi.UpdateModeFunc is nil but DefaultApi.UpdateMode was called.")
	}

	call := struct {
		Ctx     context.Context
		Subject string
	}{
		Ctx:     ctx,
		Subject: subject,
	}

	m.calls.UpdateMode = append(m.calls.UpdateMode, call)

	return m.UpdateModeFunc(ctx, subject)
}

// UpdateModeCalled returns true if UpdateMode was called at least once.
func (m *DefaultApi) UpdateModeCalled() bool {
	m.lockUpdateMode.Lock()
	defer m.lockUpdateMode.Unlock()

	return len(m.calls.UpdateMode) > 0
}

// UpdateModeCalls returns the calls made to UpdateMode.
func (m *DefaultApi) UpdateModeCalls() []struct {
	Ctx     context.Context
	Subject string
} {
	m.lockUpdateMode.Lock()
	defer m.lockUpdateMode.Unlock()

	return m.calls.UpdateMode
}

// UpdateSubjectLevelConfig mocks base method by wrapping the associated func.
func (m *DefaultApi) UpdateSubjectLevelConfig(ctx context.Context, subject string, body github_com_confluentinc_schema_registry_sdk_go.ConfigUpdateRequest) (github_com_confluentinc_schema_registry_sdk_go.ConfigUpdateRequest, *net_http.Response, error) {
	m.lockUpdateSubjectLevelConfig.Lock()
	defer m.lockUpdateSubjectLevelConfig.Unlock()

	if m.UpdateSubjectLevelConfigFunc == nil {
		panic("mocker: DefaultApi.UpdateSubjectLevelConfigFunc is nil but DefaultApi.UpdateSubjectLevelConfig was called.")
	}

	call := struct {
		Ctx     context.Context
		Subject string
		Body    github_com_confluentinc_schema_registry_sdk_go.ConfigUpdateRequest
	}{
		Ctx:     ctx,
		Subject: subject,
		Body:    body,
	}

	m.calls.UpdateSubjectLevelConfig = append(m.calls.UpdateSubjectLevelConfig, call)

	return m.UpdateSubjectLevelConfigFunc(ctx, subject, body)
}

// UpdateSubjectLevelConfigCalled returns true if UpdateSubjectLevelConfig was called at least once.
func (m *DefaultApi) UpdateSubjectLevelConfigCalled() bool {
	m.lockUpdateSubjectLevelConfig.Lock()
	defer m.lockUpdateSubjectLevelConfig.Unlock()

	return len(m.calls.UpdateSubjectLevelConfig) > 0
}

// UpdateSubjectLevelConfigCalls returns the calls made to UpdateSubjectLevelConfig.
func (m *DefaultApi) UpdateSubjectLevelConfigCalls() []struct {
	Ctx     context.Context
	Subject string
	Body    github_com_confluentinc_schema_registry_sdk_go.ConfigUpdateRequest
} {
	m.lockUpdateSubjectLevelConfig.Lock()
	defer m.lockUpdateSubjectLevelConfig.Unlock()

	return m.calls.UpdateSubjectLevelConfig
}

// UpdateTopLevelConfig mocks base method by wrapping the associated func.
func (m *DefaultApi) UpdateTopLevelConfig(ctx context.Context, body github_com_confluentinc_schema_registry_sdk_go.ConfigUpdateRequest) (github_com_confluentinc_schema_registry_sdk_go.ConfigUpdateRequest, *net_http.Response, error) {
	m.lockUpdateTopLevelConfig.Lock()
	defer m.lockUpdateTopLevelConfig.Unlock()

	if m.UpdateTopLevelConfigFunc == nil {
		panic("mocker: DefaultApi.UpdateTopLevelConfigFunc is nil but DefaultApi.UpdateTopLevelConfig was called.")
	}

	call := struct {
		Ctx  context.Context
		Body github_com_confluentinc_schema_registry_sdk_go.ConfigUpdateRequest
	}{
		Ctx:  ctx,
		Body: body,
	}

	m.calls.UpdateTopLevelConfig = append(m.calls.UpdateTopLevelConfig, call)

	return m.UpdateTopLevelConfigFunc(ctx, body)
}

// UpdateTopLevelConfigCalled returns true if UpdateTopLevelConfig was called at least once.
func (m *DefaultApi) UpdateTopLevelConfigCalled() bool {
	m.lockUpdateTopLevelConfig.Lock()
	defer m.lockUpdateTopLevelConfig.Unlock()

	return len(m.calls.UpdateTopLevelConfig) > 0
}

// UpdateTopLevelConfigCalls returns the calls made to UpdateTopLevelConfig.
func (m *DefaultApi) UpdateTopLevelConfigCalls() []struct {
	Ctx  context.Context
	Body github_com_confluentinc_schema_registry_sdk_go.ConfigUpdateRequest
} {
	m.lockUpdateTopLevelConfig.Lock()
	defer m.lockUpdateTopLevelConfig.Unlock()

	return m.calls.UpdateTopLevelConfig
}

// UpdateTopLevelMode mocks base method by wrapping the associated func.
func (m *DefaultApi) UpdateTopLevelMode(ctx context.Context) (github_com_confluentinc_schema_registry_sdk_go.ModeUpdateRequest, *net_http.Response, error) {
	m.lockUpdateTopLevelMode.Lock()
	defer m.lockUpdateTopLevelMode.Unlock()

	if m.UpdateTopLevelModeFunc == nil {
		panic("mocker: DefaultApi.UpdateTopLevelModeFunc is nil but DefaultApi.UpdateTopLevelMode was called.")
	}

	call := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}

	m.calls.UpdateTopLevelMode = append(m.calls.UpdateTopLevelMode, call)

	return m.UpdateTopLevelModeFunc(ctx)
}

// UpdateTopLevelModeCalled returns true if UpdateTopLevelMode was called at least once.
func (m *DefaultApi) UpdateTopLevelModeCalled() bool {
	m.lockUpdateTopLevelMode.Lock()
	defer m.lockUpdateTopLevelMode.Unlock()

	return len(m.calls.UpdateTopLevelMode) > 0
}

// UpdateTopLevelModeCalls returns the calls made to UpdateTopLevelMode.
func (m *DefaultApi) UpdateTopLevelModeCalls() []struct {
	Ctx context.Context
} {
	m.lockUpdateTopLevelMode.Lock()
	defer m.lockUpdateTopLevelMode.Unlock()

	return m.calls.UpdateTopLevelMode
}

// Reset resets the calls made to the mocked methods.
func (m *DefaultApi) Reset() {
	m.lockDeleteSchemaVersion.Lock()
	m.calls.DeleteSchemaVersion = nil
	m.lockDeleteSchemaVersion.Unlock()
	m.lockDeleteSubject.Lock()
	m.calls.DeleteSubject = nil
	m.lockDeleteSubject.Unlock()
	m.lockGet.Lock()
	m.calls.Get = nil
	m.lockGet.Unlock()
	m.lockGetMode.Lock()
	m.calls.GetMode = nil
	m.lockGetMode.Unlock()
	m.lockGetSchema.Lock()
	m.calls.GetSchema = nil
	m.lockGetSchema.Unlock()
	m.lockGetSchemaByVersion.Lock()
	m.calls.GetSchemaByVersion = nil
	m.lockGetSchemaByVersion.Unlock()
	m.lockGetSchemaOnly.Lock()
	m.calls.GetSchemaOnly = nil
	m.lockGetSchemaOnly.Unlock()
	m.lockGetSubjectLevelConfig.Lock()
	m.calls.GetSubjectLevelConfig = nil
	m.lockGetSubjectLevelConfig.Unlock()
	m.lockGetTopLevelConfig.Lock()
	m.calls.GetTopLevelConfig = nil
	m.lockGetTopLevelConfig.Unlock()
	m.lockGetTopLevelMode.Lock()
	m.calls.GetTopLevelMode = nil
	m.lockGetTopLevelMode.Unlock()
	m.lockList.Lock()
	m.calls.List = nil
	m.lockList.Unlock()
	m.lockListVersions.Lock()
	m.calls.ListVersions = nil
	m.lockListVersions.Unlock()
	m.lockLookUpSchemaUnderSubject.Lock()
	m.calls.LookUpSchemaUnderSubject = nil
	m.lockLookUpSchemaUnderSubject.Unlock()
	m.lockPost.Lock()
	m.calls.Post = nil
	m.lockPost.Unlock()
	m.lockRegister.Lock()
	m.calls.Register = nil
	m.lockRegister.Unlock()
	m.lockTestCompatabilityBySubjectName.Lock()
	m.calls.TestCompatabilityBySubjectName = nil
	m.lockTestCompatabilityBySubjectName.Unlock()
	m.lockUpdateMode.Lock()
	m.calls.UpdateMode = nil
	m.lockUpdateMode.Unlock()
	m.lockUpdateSubjectLevelConfig.Lock()
	m.calls.UpdateSubjectLevelConfig = nil
	m.lockUpdateSubjectLevelConfig.Unlock()
	m.lockUpdateTopLevelConfig.Lock()
	m.calls.UpdateTopLevelConfig = nil
	m.lockUpdateTopLevelConfig.Unlock()
	m.lockUpdateTopLevelMode.Lock()
	m.calls.UpdateTopLevelMode = nil
	m.lockUpdateTopLevelMode.Unlock()
}
