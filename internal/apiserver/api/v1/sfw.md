
type UserHandler struct {
	srv   srvv1.Service 
	store store.Factory
}

type service struct {
	store store.Factory
}

type Service interface {
	Users() UserSrv
	Secrets() SecretSrv
	Policies() PolicySrv
}

type Factory interface {
	Secrets() SecretStore
	Users() UserStore
	Policies() PolicyStore
	Close() error
}

// Service

type UserSrv interface {
	Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error
	Update(ctx context.Context, user *v1.User, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, usernames []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error)
	ListWithBadPerformance(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error)
	ChangePassword(ctx context.Context, user *v1.User) error
}

type SecretSrv interface {
	Create(ctx context.Context, secret *v1.Secret, opts metav1.CreateOptions) error
	Update(ctx context.Context, secret *v1.Secret, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username, secretID string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, username string, secretIDs []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username, secretID string, opts metav1.GetOptions) (*v1.Secret, error)
	List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.SecretList, error)
}

type PolicySrv interface {
	Create(ctx context.Context, policy *v1.Policy, opts metav1.CreateOptions) error
	Update(ctx context.Context, policy *v1.Policy, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username string, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, username string, names []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username string, name string, opts metav1.GetOptions) (*v1.Policy, error)
	List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.PolicyList, error)
}

type userService struct {
	store store.Factory
}

type secretService struct {
	store store.Factory
}

type policyService struct {
	store store.Factory

// Factory

type UserStore interface {
	Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error
	Update(ctx context.Context, user *v1.User, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, usernames []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error)
}

type SecretStore interface {
	Create(ctx context.Context, secret *v1.Secret, opts metav1.CreateOptions) error
	Update(ctx context.Context, secret *v1.Secret, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username, secretID string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, username string, secretIDs []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username, secretID string, opts metav1.GetOptions) (*v1.Secret, error)
	List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.SecretList, error)
}

type PolicyStore interface {
	Create(ctx context.Context, policy *v1.Policy, opts metav1.CreateOptions) error
	Update(ctx context.Context, policy *v1.Policy, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username string, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, username string, names []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username string, name string, opts metav1.GetOptions) (*v1.Policy, error)
	List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.PolicyList, error)
}

type datastore struct {
	db *gorm.DB
}

type datastore struct {
	cli             *clientv3.Client
	requestTimeout  time.Duration
	leaseTTLTimeout int

	leaseID            clientv3.LeaseID
	onKeepaliveFailure func()
	leaseLiving        bool

	watchers  map[string]*EtcdWatcher
	namespace string
}
