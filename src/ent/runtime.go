// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/sky0621/cv-admin/src/ent/careerskill"
	"github.com/sky0621/cv-admin/src/ent/careerskillgroup"
	"github.com/sky0621/cv-admin/src/ent/careertask"
	"github.com/sky0621/cv-admin/src/ent/careertaskdescription"
	"github.com/sky0621/cv-admin/src/ent/schema"
	"github.com/sky0621/cv-admin/src/ent/user"
	"github.com/sky0621/cv-admin/src/ent/useractivity"
	"github.com/sky0621/cv-admin/src/ent/usercareer"
	"github.com/sky0621/cv-admin/src/ent/usercareerdescription"
	"github.com/sky0621/cv-admin/src/ent/usercareergroup"
	"github.com/sky0621/cv-admin/src/ent/usernote"
	"github.com/sky0621/cv-admin/src/ent/usernoteitem"
	"github.com/sky0621/cv-admin/src/ent/userqualification"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	careerskillMixin := schema.CareerSkill{}.Mixin()
	careerskillMixinFields0 := careerskillMixin[0].Fields()
	_ = careerskillMixinFields0
	careerskillFields := schema.CareerSkill{}.Fields()
	_ = careerskillFields
	// careerskillDescCreateTime is the schema descriptor for create_time field.
	careerskillDescCreateTime := careerskillMixinFields0[0].Descriptor()
	// careerskill.DefaultCreateTime holds the default value on creation for the create_time field.
	careerskill.DefaultCreateTime = careerskillDescCreateTime.Default.(func() time.Time)
	// careerskillDescUpdateTime is the schema descriptor for update_time field.
	careerskillDescUpdateTime := careerskillMixinFields0[1].Descriptor()
	// careerskill.DefaultUpdateTime holds the default value on creation for the update_time field.
	careerskill.DefaultUpdateTime = careerskillDescUpdateTime.Default.(func() time.Time)
	// careerskill.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	careerskill.UpdateDefaultUpdateTime = careerskillDescUpdateTime.UpdateDefault.(func() time.Time)
	// careerskillDescName is the schema descriptor for name field.
	careerskillDescName := careerskillFields[0].Descriptor()
	// careerskill.NameValidator is a validator for the "name" field. It is called by the builders before save.
	careerskill.NameValidator = func() func(string) error {
		validators := careerskillDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// careerskillDescURL is the schema descriptor for url field.
	careerskillDescURL := careerskillFields[1].Descriptor()
	// careerskill.URLValidator is a validator for the "url" field. It is called by the builders before save.
	careerskill.URLValidator = careerskillDescURL.Validators[0].(func(string) error)
	// careerskillDescVersion is the schema descriptor for version field.
	careerskillDescVersion := careerskillFields[2].Descriptor()
	// careerskill.VersionValidator is a validator for the "version" field. It is called by the builders before save.
	careerskill.VersionValidator = careerskillDescVersion.Validators[0].(func(string) error)
	careerskillgroupMixin := schema.CareerSkillGroup{}.Mixin()
	careerskillgroupMixinFields0 := careerskillgroupMixin[0].Fields()
	_ = careerskillgroupMixinFields0
	careerskillgroupFields := schema.CareerSkillGroup{}.Fields()
	_ = careerskillgroupFields
	// careerskillgroupDescCreateTime is the schema descriptor for create_time field.
	careerskillgroupDescCreateTime := careerskillgroupMixinFields0[0].Descriptor()
	// careerskillgroup.DefaultCreateTime holds the default value on creation for the create_time field.
	careerskillgroup.DefaultCreateTime = careerskillgroupDescCreateTime.Default.(func() time.Time)
	// careerskillgroupDescUpdateTime is the schema descriptor for update_time field.
	careerskillgroupDescUpdateTime := careerskillgroupMixinFields0[1].Descriptor()
	// careerskillgroup.DefaultUpdateTime holds the default value on creation for the update_time field.
	careerskillgroup.DefaultUpdateTime = careerskillgroupDescUpdateTime.Default.(func() time.Time)
	// careerskillgroup.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	careerskillgroup.UpdateDefaultUpdateTime = careerskillgroupDescUpdateTime.UpdateDefault.(func() time.Time)
	// careerskillgroupDescLabel is the schema descriptor for label field.
	careerskillgroupDescLabel := careerskillgroupFields[0].Descriptor()
	// careerskillgroup.LabelValidator is a validator for the "label" field. It is called by the builders before save.
	careerskillgroup.LabelValidator = func() func(string) error {
		validators := careerskillgroupDescLabel.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(label string) error {
			for _, fn := range fns {
				if err := fn(label); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	careertaskMixin := schema.CareerTask{}.Mixin()
	careertaskMixinFields0 := careertaskMixin[0].Fields()
	_ = careertaskMixinFields0
	careertaskFields := schema.CareerTask{}.Fields()
	_ = careertaskFields
	// careertaskDescCreateTime is the schema descriptor for create_time field.
	careertaskDescCreateTime := careertaskMixinFields0[0].Descriptor()
	// careertask.DefaultCreateTime holds the default value on creation for the create_time field.
	careertask.DefaultCreateTime = careertaskDescCreateTime.Default.(func() time.Time)
	// careertaskDescUpdateTime is the schema descriptor for update_time field.
	careertaskDescUpdateTime := careertaskMixinFields0[1].Descriptor()
	// careertask.DefaultUpdateTime holds the default value on creation for the update_time field.
	careertask.DefaultUpdateTime = careertaskDescUpdateTime.Default.(func() time.Time)
	// careertask.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	careertask.UpdateDefaultUpdateTime = careertaskDescUpdateTime.UpdateDefault.(func() time.Time)
	// careertaskDescName is the schema descriptor for name field.
	careertaskDescName := careertaskFields[0].Descriptor()
	// careertask.NameValidator is a validator for the "name" field. It is called by the builders before save.
	careertask.NameValidator = func() func(string) error {
		validators := careertaskDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	careertaskdescriptionFields := schema.CareerTaskDescription{}.Fields()
	_ = careertaskdescriptionFields
	// careertaskdescriptionDescDescription is the schema descriptor for description field.
	careertaskdescriptionDescDescription := careertaskdescriptionFields[0].Descriptor()
	// careertaskdescription.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	careertaskdescription.DescriptionValidator = func() func(string) error {
		validators := careertaskdescriptionDescDescription.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(description string) error {
			for _, fn := range fns {
				if err := fn(description); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = func() func(string) error {
		validators := userDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescNickname is the schema descriptor for nickname field.
	userDescNickname := userFields[1].Descriptor()
	// user.NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	user.NicknameValidator = userDescNickname.Validators[0].(func(string) error)
	// userDescAvatarURL is the schema descriptor for avatar_url field.
	userDescAvatarURL := userFields[2].Descriptor()
	// user.AvatarURLValidator is a validator for the "avatar_url" field. It is called by the builders before save.
	user.AvatarURLValidator = userDescAvatarURL.Validators[0].(func(string) error)
	// userDescBirthdayYear is the schema descriptor for birthday_year field.
	userDescBirthdayYear := userFields[3].Descriptor()
	// user.BirthdayYearValidator is a validator for the "birthday_year" field. It is called by the builders before save.
	user.BirthdayYearValidator = userDescBirthdayYear.Validators[0].(func(int) error)
	// userDescBirthdayMonth is the schema descriptor for birthday_month field.
	userDescBirthdayMonth := userFields[4].Descriptor()
	// user.BirthdayMonthValidator is a validator for the "birthday_month" field. It is called by the builders before save.
	user.BirthdayMonthValidator = userDescBirthdayMonth.Validators[0].(func(int) error)
	// userDescBirthdayDay is the schema descriptor for birthday_day field.
	userDescBirthdayDay := userFields[5].Descriptor()
	// user.BirthdayDayValidator is a validator for the "birthday_day" field. It is called by the builders before save.
	user.BirthdayDayValidator = userDescBirthdayDay.Validators[0].(func(int) error)
	// userDescJob is the schema descriptor for job field.
	userDescJob := userFields[6].Descriptor()
	// user.JobValidator is a validator for the "job" field. It is called by the builders before save.
	user.JobValidator = userDescJob.Validators[0].(func(string) error)
	// userDescBelongTo is the schema descriptor for belong_to field.
	userDescBelongTo := userFields[7].Descriptor()
	// user.BelongToValidator is a validator for the "belong_to" field. It is called by the builders before save.
	user.BelongToValidator = userDescBelongTo.Validators[0].(func(string) error)
	// userDescPr is the schema descriptor for pr field.
	userDescPr := userFields[8].Descriptor()
	// user.PrValidator is a validator for the "pr" field. It is called by the builders before save.
	user.PrValidator = userDescPr.Validators[0].(func(string) error)
	useractivityMixin := schema.UserActivity{}.Mixin()
	useractivityMixinFields0 := useractivityMixin[0].Fields()
	_ = useractivityMixinFields0
	useractivityFields := schema.UserActivity{}.Fields()
	_ = useractivityFields
	// useractivityDescCreateTime is the schema descriptor for create_time field.
	useractivityDescCreateTime := useractivityMixinFields0[0].Descriptor()
	// useractivity.DefaultCreateTime holds the default value on creation for the create_time field.
	useractivity.DefaultCreateTime = useractivityDescCreateTime.Default.(func() time.Time)
	// useractivityDescUpdateTime is the schema descriptor for update_time field.
	useractivityDescUpdateTime := useractivityMixinFields0[1].Descriptor()
	// useractivity.DefaultUpdateTime holds the default value on creation for the update_time field.
	useractivity.DefaultUpdateTime = useractivityDescUpdateTime.Default.(func() time.Time)
	// useractivity.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	useractivity.UpdateDefaultUpdateTime = useractivityDescUpdateTime.UpdateDefault.(func() time.Time)
	// useractivityDescName is the schema descriptor for name field.
	useractivityDescName := useractivityFields[0].Descriptor()
	// useractivity.NameValidator is a validator for the "name" field. It is called by the builders before save.
	useractivity.NameValidator = func() func(string) error {
		validators := useractivityDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// useractivityDescURL is the schema descriptor for url field.
	useractivityDescURL := useractivityFields[1].Descriptor()
	// useractivity.URLValidator is a validator for the "url" field. It is called by the builders before save.
	useractivity.URLValidator = useractivityDescURL.Validators[0].(func(string) error)
	// useractivityDescIcon is the schema descriptor for icon field.
	useractivityDescIcon := useractivityFields[2].Descriptor()
	// useractivity.IconValidator is a validator for the "icon" field. It is called by the builders before save.
	useractivity.IconValidator = useractivityDescIcon.Validators[0].(func(string) error)
	usercareerMixin := schema.UserCareer{}.Mixin()
	usercareerMixinFields0 := usercareerMixin[0].Fields()
	_ = usercareerMixinFields0
	usercareerFields := schema.UserCareer{}.Fields()
	_ = usercareerFields
	// usercareerDescCreateTime is the schema descriptor for create_time field.
	usercareerDescCreateTime := usercareerMixinFields0[0].Descriptor()
	// usercareer.DefaultCreateTime holds the default value on creation for the create_time field.
	usercareer.DefaultCreateTime = usercareerDescCreateTime.Default.(func() time.Time)
	// usercareerDescUpdateTime is the schema descriptor for update_time field.
	usercareerDescUpdateTime := usercareerMixinFields0[1].Descriptor()
	// usercareer.DefaultUpdateTime holds the default value on creation for the update_time field.
	usercareer.DefaultUpdateTime = usercareerDescUpdateTime.Default.(func() time.Time)
	// usercareer.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	usercareer.UpdateDefaultUpdateTime = usercareerDescUpdateTime.UpdateDefault.(func() time.Time)
	// usercareerDescName is the schema descriptor for name field.
	usercareerDescName := usercareerFields[0].Descriptor()
	// usercareer.NameValidator is a validator for the "name" field. It is called by the builders before save.
	usercareer.NameValidator = func() func(string) error {
		validators := usercareerDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// usercareerDescFrom is the schema descriptor for from field.
	usercareerDescFrom := usercareerFields[1].Descriptor()
	// usercareer.FromValidator is a validator for the "from" field. It is called by the builders before save.
	usercareer.FromValidator = func() func(string) error {
		validators := usercareerDescFrom.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(from string) error {
			for _, fn := range fns {
				if err := fn(from); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// usercareerDescTo is the schema descriptor for to field.
	usercareerDescTo := usercareerFields[2].Descriptor()
	// usercareer.ToValidator is a validator for the "to" field. It is called by the builders before save.
	usercareer.ToValidator = usercareerDescTo.Validators[0].(func(string) error)
	usercareerdescriptionFields := schema.UserCareerDescription{}.Fields()
	_ = usercareerdescriptionFields
	// usercareerdescriptionDescDescription is the schema descriptor for description field.
	usercareerdescriptionDescDescription := usercareerdescriptionFields[0].Descriptor()
	// usercareerdescription.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	usercareerdescription.DescriptionValidator = func() func(string) error {
		validators := usercareerdescriptionDescDescription.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(description string) error {
			for _, fn := range fns {
				if err := fn(description); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	usercareergroupMixin := schema.UserCareerGroup{}.Mixin()
	usercareergroupMixinFields0 := usercareergroupMixin[0].Fields()
	_ = usercareergroupMixinFields0
	usercareergroupFields := schema.UserCareerGroup{}.Fields()
	_ = usercareergroupFields
	// usercareergroupDescCreateTime is the schema descriptor for create_time field.
	usercareergroupDescCreateTime := usercareergroupMixinFields0[0].Descriptor()
	// usercareergroup.DefaultCreateTime holds the default value on creation for the create_time field.
	usercareergroup.DefaultCreateTime = usercareergroupDescCreateTime.Default.(func() time.Time)
	// usercareergroupDescUpdateTime is the schema descriptor for update_time field.
	usercareergroupDescUpdateTime := usercareergroupMixinFields0[1].Descriptor()
	// usercareergroup.DefaultUpdateTime holds the default value on creation for the update_time field.
	usercareergroup.DefaultUpdateTime = usercareergroupDescUpdateTime.Default.(func() time.Time)
	// usercareergroup.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	usercareergroup.UpdateDefaultUpdateTime = usercareergroupDescUpdateTime.UpdateDefault.(func() time.Time)
	// usercareergroupDescLabel is the schema descriptor for label field.
	usercareergroupDescLabel := usercareergroupFields[0].Descriptor()
	// usercareergroup.LabelValidator is a validator for the "label" field. It is called by the builders before save.
	usercareergroup.LabelValidator = func() func(string) error {
		validators := usercareergroupDescLabel.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(label string) error {
			for _, fn := range fns {
				if err := fn(label); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	usernoteMixin := schema.UserNote{}.Mixin()
	usernoteMixinFields0 := usernoteMixin[0].Fields()
	_ = usernoteMixinFields0
	usernoteFields := schema.UserNote{}.Fields()
	_ = usernoteFields
	// usernoteDescCreateTime is the schema descriptor for create_time field.
	usernoteDescCreateTime := usernoteMixinFields0[0].Descriptor()
	// usernote.DefaultCreateTime holds the default value on creation for the create_time field.
	usernote.DefaultCreateTime = usernoteDescCreateTime.Default.(func() time.Time)
	// usernoteDescUpdateTime is the schema descriptor for update_time field.
	usernoteDescUpdateTime := usernoteMixinFields0[1].Descriptor()
	// usernote.DefaultUpdateTime holds the default value on creation for the update_time field.
	usernote.DefaultUpdateTime = usernoteDescUpdateTime.Default.(func() time.Time)
	// usernote.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	usernote.UpdateDefaultUpdateTime = usernoteDescUpdateTime.UpdateDefault.(func() time.Time)
	// usernoteDescLabel is the schema descriptor for label field.
	usernoteDescLabel := usernoteFields[0].Descriptor()
	// usernote.LabelValidator is a validator for the "label" field. It is called by the builders before save.
	usernote.LabelValidator = func() func(string) error {
		validators := usernoteDescLabel.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(label string) error {
			for _, fn := range fns {
				if err := fn(label); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// usernoteDescMemo is the schema descriptor for memo field.
	usernoteDescMemo := usernoteFields[1].Descriptor()
	// usernote.MemoValidator is a validator for the "memo" field. It is called by the builders before save.
	usernote.MemoValidator = usernoteDescMemo.Validators[0].(func(string) error)
	usernoteitemMixin := schema.UserNoteItem{}.Mixin()
	usernoteitemMixinFields0 := usernoteitemMixin[0].Fields()
	_ = usernoteitemMixinFields0
	usernoteitemFields := schema.UserNoteItem{}.Fields()
	_ = usernoteitemFields
	// usernoteitemDescCreateTime is the schema descriptor for create_time field.
	usernoteitemDescCreateTime := usernoteitemMixinFields0[0].Descriptor()
	// usernoteitem.DefaultCreateTime holds the default value on creation for the create_time field.
	usernoteitem.DefaultCreateTime = usernoteitemDescCreateTime.Default.(func() time.Time)
	// usernoteitemDescUpdateTime is the schema descriptor for update_time field.
	usernoteitemDescUpdateTime := usernoteitemMixinFields0[1].Descriptor()
	// usernoteitem.DefaultUpdateTime holds the default value on creation for the update_time field.
	usernoteitem.DefaultUpdateTime = usernoteitemDescUpdateTime.Default.(func() time.Time)
	// usernoteitem.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	usernoteitem.UpdateDefaultUpdateTime = usernoteitemDescUpdateTime.UpdateDefault.(func() time.Time)
	// usernoteitemDescText is the schema descriptor for text field.
	usernoteitemDescText := usernoteitemFields[0].Descriptor()
	// usernoteitem.TextValidator is a validator for the "text" field. It is called by the builders before save.
	usernoteitem.TextValidator = func() func(string) error {
		validators := usernoteitemDescText.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(text string) error {
			for _, fn := range fns {
				if err := fn(text); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	userqualificationMixin := schema.UserQualification{}.Mixin()
	userqualificationMixinFields0 := userqualificationMixin[0].Fields()
	_ = userqualificationMixinFields0
	userqualificationFields := schema.UserQualification{}.Fields()
	_ = userqualificationFields
	// userqualificationDescCreateTime is the schema descriptor for create_time field.
	userqualificationDescCreateTime := userqualificationMixinFields0[0].Descriptor()
	// userqualification.DefaultCreateTime holds the default value on creation for the create_time field.
	userqualification.DefaultCreateTime = userqualificationDescCreateTime.Default.(func() time.Time)
	// userqualificationDescUpdateTime is the schema descriptor for update_time field.
	userqualificationDescUpdateTime := userqualificationMixinFields0[1].Descriptor()
	// userqualification.DefaultUpdateTime holds the default value on creation for the update_time field.
	userqualification.DefaultUpdateTime = userqualificationDescUpdateTime.Default.(func() time.Time)
	// userqualification.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	userqualification.UpdateDefaultUpdateTime = userqualificationDescUpdateTime.UpdateDefault.(func() time.Time)
	// userqualificationDescName is the schema descriptor for name field.
	userqualificationDescName := userqualificationFields[0].Descriptor()
	// userqualification.NameValidator is a validator for the "name" field. It is called by the builders before save.
	userqualification.NameValidator = func() func(string) error {
		validators := userqualificationDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userqualificationDescOrganization is the schema descriptor for organization field.
	userqualificationDescOrganization := userqualificationFields[1].Descriptor()
	// userqualification.OrganizationValidator is a validator for the "organization" field. It is called by the builders before save.
	userqualification.OrganizationValidator = userqualificationDescOrganization.Validators[0].(func(string) error)
	// userqualificationDescURL is the schema descriptor for url field.
	userqualificationDescURL := userqualificationFields[2].Descriptor()
	// userqualification.URLValidator is a validator for the "url" field. It is called by the builders before save.
	userqualification.URLValidator = userqualificationDescURL.Validators[0].(func(string) error)
	// userqualificationDescGotDate is the schema descriptor for got_date field.
	userqualificationDescGotDate := userqualificationFields[3].Descriptor()
	// userqualification.GotDateValidator is a validator for the "got_date" field. It is called by the builders before save.
	userqualification.GotDateValidator = userqualificationDescGotDate.Validators[0].(func(string) error)
	// userqualificationDescMemo is the schema descriptor for memo field.
	userqualificationDescMemo := userqualificationFields[4].Descriptor()
	// userqualification.MemoValidator is a validator for the "memo" field. It is called by the builders before save.
	userqualification.MemoValidator = userqualificationDescMemo.Validators[0].(func(string) error)
}
