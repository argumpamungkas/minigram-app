import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:minigram/config/constant.dart';
import 'package:minigram/view/auth/widgets/text_form_field_custom.dart';

class LoginScreen extends StatelessWidget {
  const LoginScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: SingleChildScrollView(
          child: Padding(
            padding: EdgeInsets.only(top: 0.1.sh),
            child: Column(
              spacing: 8.r,
              children: [
                Image.asset(Constant.imgLogo, width: 0.2.sw),
                Form(
                  child: Card(
                    color: Colors.white,
                    margin: EdgeInsets.all(24.r),
                    elevation: 4,
                    child: Padding(
                      padding: EdgeInsets.all(16.r),
                      child: Column(
                        spacing: 12.r,
                        children: [
                          Text(
                            "LOGIN",
                            style: TextStyle(
                              fontSize: 20.r,
                              letterSpacing: 4.r,
                              fontWeight: FontWeight.bold,
                            ),
                          ),
                          TextFormFieldCustom(
                            hintText: 'username',
                            icon: Icons.person,
                          ),
                          Column(
                            children: [
                              TextFormFieldCustom(
                                hintText: 'password',
                                icon: Icons.lock,
                                suffix: Icon(Icons.visibility_off),
                              ),
                              Align(
                                alignment: Alignment.centerRight,
                                child: TextButton(
                                  onPressed: () {},
                                  child: Text("Forgot Password?"),
                                ),
                              ),
                            ],
                          ),
                          ElevatedButton(
                            onPressed: () {},
                            child: Text("Login"),
                          ),
                          Row(
                            mainAxisAlignment: MainAxisAlignment.center,
                            children: [
                              Text("Haven't an account?"),
                              TextButton(
                                style: TextButton.styleFrom(
                                  padding: EdgeInsets.zero,
                                ),
                                onPressed: () {},
                                child: Text("Register"),
                              ),
                            ],
                          ),
                        ],
                      ),
                    ),
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
