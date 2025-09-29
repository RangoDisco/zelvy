package com.zelvy.companion

import android.app.Application
import com.zelvy.companion.data.HealthConnectManager

class BaseApplication: Application() {
    val healthConnectManager by lazy {
        HealthConnectManager(this)
    }
}