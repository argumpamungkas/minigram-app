import 'package:animated_splash_screen/animated_splash_screen.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:minigram/config/constant.dart';
import 'package:minigram/view/auth/login_screen.dart';

class SplashScreen extends StatelessWidget {
  const SplashScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: AnimatedSplashScreen(
        splash: Constant.imgLogo,
        centered: true,
        splashTransition: SplashTransition.scaleTransition,
        splashIconSize: 0.3.sw,
        nextScreen: LoginScreen(),
      ),
    );
  }
}
