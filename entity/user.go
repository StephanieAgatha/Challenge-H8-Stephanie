package entity

import (
	"fmt"
	"gorm.io/gorm"
)

// https://gorm.io/docs/models.html
type User struct {
	*gorm.Model

	Username string `db:"username"`
	Password string `db:"password"`
}

func (u *User) PrintSuccess() {

	//print welcome
	fmt.Println("Selamat Datang", u.Username)
}

//gorm model adalah tipe data yang diperluas oleh gorm untuk menyediakan fitur fitur tambahan yang terkait dengan ORM.
//tipe *gorm.model digunakan sebagai jenis dasar atau parent model yang digunakan dalam definisi model data

//dalam struktur user.*gorm model digunakan untuk memperluas fitur fitur seperti penanda waktu,pengaturan referensi primary key,
//dan pengaturan referensi terhadap tabel di dalam basis data.
//secara khusus fitur fitur yang disediakan oleh *gorm.model adalah sebagai berikut
//1.ID : menyediakan kolom id sebagai primary key
//2.CreatedAt : menyediakan kolom dengan nama "created_at" yang akan secara otomatis diisi dengan waktu saat dibuat.
//3.UpdateAt : Menyediakan kolom dengan nama "updated_at" yang akan secara otomatis diperbarui dengan waktu saat data diperbarui
//4.DeletedAt : menyediakan kolom dengan nama "deleted_at" yang akan secara otomatis diisi dengan waktu saat data dihapus
//(jika fitur "soft delete" diaktifkan).

//dengan menggunakan *gorm.model sebagai tipe data yang diperluan oleh struktur User,kita dapat menggunakan fitur-fitur ini secara langsung
//dalam operasi-operasi ORM dengan GORM tanpa perlu mendefinisikannya secara manual
